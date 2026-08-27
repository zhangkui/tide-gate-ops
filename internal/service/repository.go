package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"

	"gitlab.com/zhangkui/tide-gate-ops/internal/clock"
	"gitlab.com/zhangkui/tide-gate-ops/internal/model"
	"gitlab.com/zhangkui/tide-gate-ops/internal/store"
)

const (
	kindStation     = "station"
	kindGate        = "gate"
	kindSensor      = "sensor"
	kindReading     = "reading"
	kindWindow      = "window"
	kindCommand     = "command"
	kindAlarm       = "alarm"
	kindShift       = "shift"
	kindHandover    = "handover"
	kindMaintenance = "maintenance"
)

type Lab struct {
	store   *store.Store
	clock   *clock.Clock
	mu      sync.RWMutex
	closed  chan struct{}
	workers sync.WaitGroup
}

func NewLab(s *store.Store) *Lab {
	return &Lab{store: s, clock: clock.New(), closed: make(chan struct{})}
}

func (l *Lab) Close() {
	select {
	case <-l.closed:
		return
	default:
		close(l.closed)
	}
	l.workers.Wait()
}
func (l *Lab) Clock() *clock.Clock { return l.clock }
func (l *Lab) Store() *store.Store { return l.store }

func save[T any](ctx context.Context, s *store.Store, kind, id string, value T) error {
	return s.SaveContext(ctx, kind, id, value)
}
func load[T any](ctx context.Context, s *store.Store, kind, id string) (T, error) {
	var value T
	err := s.LoadContext(ctx, kind, id, &value)
	return value, err
}
func list[T any](ctx context.Context, s *store.Store, kind string) ([]T, error) {
	out := make([]T, 0)
	err := s.ListContext(ctx, kind, func(raw []byte) error {
		var v T
		if err := model.Unmarshal(raw, &v); err != nil {
			return err
		}
		out = append(out, v)
		return nil
	})
	return out, err
}

func (l *Lab) CreateStation(ctx context.Context, station model.TideStation) error {
	if err := station.Validate(); err != nil {
		return err
	}
	station.CreatedAt = l.clock.Now()
	if err := save(ctx, l.store, kindStation, station.ID, station); err != nil {
		return fmt.Errorf("create station: %w", err)
	}
	return l.store.AppendEvent(ctx, station.ID, "station.created", station)
}

func (l *Lab) GetStation(ctx context.Context, id string) (model.TideStation, error) {
	value, err := load[model.TideStation](ctx, l.store, kindStation, id)
	if errors.Is(err, sql.ErrNoRows) {
		return value, model.ErrNotFound
	}
	return value, err
}
func (l *Lab) ListStations(ctx context.Context) ([]model.TideStation, error) {
	return list[model.TideStation](ctx, l.store, kindStation)
}

func (l *Lab) AddGate(ctx context.Context, gate model.Gate) error {
	if err := gate.Validate(); err != nil {
		return err
	}
	if _, err := l.GetStation(ctx, gate.StationID); err != nil {
		return fmt.Errorf("gate station: %w", err)
	}
	gate.UpdatedAt = l.clock.Now()
	if gate.Status == "" {
		gate.Status = model.GateClosed
	}
	if err := save(ctx, l.store, kindGate, gate.ID, gate); err != nil {
		return fmt.Errorf("add gate: %w", err)
	}
	return l.store.AppendEvent(ctx, gate.ID, "gate.created", gate)
}

func (l *Lab) GetGate(ctx context.Context, id string) (model.Gate, error) {
	value, err := load[model.Gate](ctx, l.store, kindGate, id)
	if errors.Is(err, sql.ErrNoRows) {
		return value, model.ErrNotFound
	}
	return value, err
}
func (l *Lab) ListGates(ctx context.Context, stationID string) ([]model.Gate, error) {
	all, err := list[model.Gate](ctx, l.store, kindGate)
	if err != nil {
		return nil, err
	}
	out := all[:0]
	for _, g := range all {
		if stationID == "" || g.StationID == stationID {
			out = append(out, g)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}

func (l *Lab) RegisterSensor(ctx context.Context, sensor model.Sensor) error {
	if sensor.ID == "" || sensor.StationID == "" || sensor.Kind == "" {
		return model.ErrInvalidReading
	}
	if _, err := l.GetStation(ctx, sensor.StationID); err != nil {
		return err
	}
	if err := save(ctx, l.store, kindSensor, sensor.ID, sensor); err != nil {
		return fmt.Errorf("register sensor: %w", err)
	}
	return l.store.AppendEvent(ctx, sensor.ID, "sensor.registered", sensor)
}
func (l *Lab) GetSensor(ctx context.Context, id string) (model.Sensor, error) {
	v, e := load[model.Sensor](ctx, l.store, kindSensor, id)
	if errors.Is(e, sql.ErrNoRows) {
		return v, model.ErrNotFound
	}
	return v, e
}
func (l *Lab) ListSensors(ctx context.Context, stationID string) ([]model.Sensor, error) {
	all, e := list[model.Sensor](ctx, l.store, kindSensor)
	if e != nil {
		return nil, e
	}
	out := all[:0]
	for _, s := range all {
		if stationID == "" || s.StationID == stationID {
			out = append(out, s)
		}
	}
	return out, nil
}

func (l *Lab) IngestReading(ctx context.Context, reading model.Reading) error {
	if err := reading.Validate(); err != nil {
		return err
	}
	sensor, err := l.GetSensor(ctx, reading.SensorID)
	if err != nil {
		return fmt.Errorf("reading sensor: %w", err)
	}
	if sensor.StationID != reading.StationID || sensor.Kind != reading.Kind {
		return fmt.Errorf("%w: sensor mapping", model.ErrInvalidReading)
	}
	reading.Quality = model.NormalizeQuality(reading.Quality)
	if reading.ReceivedAt.IsZero() {
		reading.ReceivedAt = l.clock.Now()
	}
	if err := save(ctx, l.store, kindReading, reading.ID, reading); err != nil {
		return fmt.Errorf("ingest reading: %w", err)
	}
	sensor.Online = true
	sensor.LastSeen = reading.ReceivedAt
	_ = save(ctx, l.store, kindSensor, sensor.ID, sensor)
	return l.store.AppendEvent(ctx, reading.StationID, "reading.ingested", reading)
}

func (l *Lab) ListReadings(ctx context.Context, stationID, kind string, from, to time.Time) ([]model.Reading, error) {
	all, e := list[model.Reading](ctx, l.store, kindReading)
	if e != nil {
		return nil, e
	}
	out := make([]model.Reading, 0, len(all))
	for _, r := range all {
		if stationID != "" && r.StationID != stationID {
			continue
		}
		if kind != "" && r.Kind != kind {
			continue
		}
		if !from.IsZero() && r.ObservedAt.Before(from) {
			continue
		}
		if !to.IsZero() && r.ObservedAt.After(to) {
			continue
		}
		out = append(out, r)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ObservedAt.Before(out[j].ObservedAt) })
	return out, nil
}

func (l *Lab) AddTideWindow(ctx context.Context, w model.TideWindow) error {
	if w.ID == "" || w.StationID == "" || w.At.IsZero() {
		return model.ErrInvalidReading
	}
	if _, e := l.GetStation(ctx, w.StationID); e != nil {
		return e
	}
	if e := save(ctx, l.store, kindWindow, w.ID, w); e != nil {
		return e
	}
	return l.store.AppendEvent(ctx, w.StationID, "tide.window_added", w)
}
func (l *Lab) ListTideWindows(ctx context.Context, stationID string) ([]model.TideWindow, error) {
	all, e := list[model.TideWindow](ctx, l.store, kindWindow)
	if e != nil {
		return nil, e
	}
	out := all[:0]
	for _, w := range all {
		if stationID == "" || w.StationID == stationID {
			out = append(out, w)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out, nil
}

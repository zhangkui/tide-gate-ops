package service

import (
	"context"
	"errors"
	"path/filepath"
	"testing"

	"gitlab.com/zhangkui/tide-gate-ops/internal/model"
	"gitlab.com/zhangkui/tide-gate-ops/internal/store"
)

func newTestLab(t *testing.T) *Lab {
	t.Helper()
	s, err := store.Open(filepath.Join(t.TempDir(), "tide.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = s.Close() })
	return NewLab(s)
}
func seedStationSensor(t *testing.T, l *Lab) {
	t.Helper()
	ctx := context.Background()
	if err := l.CreateStation(ctx, model.TideStation{ID: "station-1", Name: "Station", Latitude: 1, Longitude: 1, Timezone: "UTC", Active: true}); err != nil {
		t.Fatal(err)
	}
	if err := l.RegisterSensor(ctx, model.Sensor{ID: "sensor-1", StationID: "station-1", Kind: "tide_level", Unit: "m", Min: -10, Max: 20}); err != nil {
		t.Fatal(err)
	}
}

// TestBug07_TideGate verifies that a cancelled context aborts the ingest
// chain. The injected defect replaces the caller context with a detached
// background context on every persistence call, so the cancel signal never
// reaches the write path and the request completes despite cancellation.
func TestBug07_TideGate(t *testing.T) {
	l := newTestLab(t)
	seedStationSensor(t, l)
	now := l.Clock().Now()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := l.IngestReading(ctx, model.Reading{ID: "r1", SensorID: "sensor-1", StationID: "station-1", Kind: "tide_level", Value: 2, Unit: "m", Quality: "good", ObservedAt: now, ReceivedAt: now})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled ingest must abort with context.Canceled, got %v", err)
	}
}
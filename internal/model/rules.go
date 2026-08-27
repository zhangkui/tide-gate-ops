package model

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
)

var (
	ErrInvalidStation  = errors.New("station is invalid")
	ErrInvalidGate     = errors.New("gate is invalid")
	ErrInvalidReading  = errors.New("reading is invalid")
	ErrInvalidCommand  = errors.New("gate command is invalid")
	ErrUnsafeOperation = errors.New("operation violates safety interlock")
	ErrStateConflict   = errors.New("gate state version conflict")
	ErrNotFound        = errors.New("entity not found")
)

func (s TideStation) Validate() error {
	if strings.TrimSpace(s.ID) == "" || strings.TrimSpace(s.Name) == "" {
		return fmt.Errorf("%w: id and name required", ErrInvalidStation)
	}
	if s.Latitude < -90 || s.Latitude > 90 || s.Longitude < -180 || s.Longitude > 180 {
		return fmt.Errorf("%w: coordinates", ErrInvalidStation)
	}
	return nil
}

func (g Gate) Validate() error {
	if strings.TrimSpace(g.ID) == "" || strings.TrimSpace(g.StationID) == "" {
		return fmt.Errorf("%w: identity", ErrInvalidGate)
	}
	if g.OpeningPercent < 0 || g.OpeningPercent > 100 || g.TargetPercent < 0 || g.TargetPercent > 100 {
		return fmt.Errorf("%w: opening range", ErrInvalidGate)
	}
	if g.MaxWindKPH < 0 || g.MaxRiseRate < 0 {
		return fmt.Errorf("%w: safety limits", ErrInvalidGate)
	}
	return nil
}

func (r Reading) Validate() error {
	if strings.TrimSpace(r.ID) == "" || strings.TrimSpace(r.SensorID) == "" || strings.TrimSpace(r.StationID) == "" {
		return fmt.Errorf("%w: identity", ErrInvalidReading)
	}
	if math.IsNaN(r.Value) || math.IsInf(r.Value, 0) || r.ObservedAt.IsZero() {
		return fmt.Errorf("%w: value or timestamp", ErrInvalidReading)
	}
	if r.ReceivedAt.IsZero() {
		return fmt.Errorf("%w: received timestamp", ErrInvalidReading)
	}
	return nil
}

func (c GateCommand) Validate() error {
	if strings.TrimSpace(c.ID) == "" || strings.TrimSpace(c.GateID) == "" || strings.TrimSpace(c.Operator) == "" {
		return fmt.Errorf("%w: identity", ErrInvalidCommand)
	}
	if c.TargetPercent < 0 || c.TargetPercent > 100 {
		return fmt.Errorf("%w: target range", ErrInvalidCommand)
	}
	return nil
}

func (g Gate) CanMove(target float64, now time.Time) error {
	if g.Status == GateFault {
		return fmt.Errorf("%w: gate in fault", ErrUnsafeOperation)
	}
	if target < 0 || target > 100 {
		return fmt.Errorf("%w: target outside range", ErrUnsafeOperation)
	}
	if now.IsZero() {
		return fmt.Errorf("%w: clock unavailable", ErrUnsafeOperation)
	}
	return nil
}

func (g Gate) Transition(target float64) GateStatus {
	if target == 0 {
		return GateClosing
	}
	if target == 100 {
		return GateOpening
	}
	if target > g.OpeningPercent {
		return GateOpening
	}
	return GateClosing
}

func NormalizeQuality(q string) string {
	q = strings.ToLower(strings.TrimSpace(q))
	if q == "" {
		return "unknown"
	}
	if q != "good" && q != "suspect" && q != "bad" {
		return "suspect"
	}
	return q
}

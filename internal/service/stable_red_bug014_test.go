package service

import (
	"context"
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

// TestBug14_TideGate verifies that a sensor constraint check evaluates the raw
// reading against the sensor's physical range, not a value that was already
// shifted by calibration. The injected defect applies calibration again inside
// the check, so an in-range raw reading is wrongly flagged as above-max.
func TestBug14_TideGate(t *testing.T) {
	l := newTestLab(t)
	seedStationSensor(t, l)
	ctx := context.Background()
	if _, e := l.CalibrateSensor(ctx, model.Calibration{ID: "cal-1", SensorID: "sensor-1", Offset: 15, Scale: 1, Operator: "op", EffectiveAt: l.Clock().Now()}); e != nil {
		t.Fatal(e)
	}
	verdict, err := l.CheckSensorConstraint01(ctx, "sensor-1", 10)
	if err != nil {
		t.Fatal(err)
	}
	if verdict != "within-range" {
		t.Fatalf("raw reading 10 with offset 15 must be within range, got %q", verdict)
	}
}
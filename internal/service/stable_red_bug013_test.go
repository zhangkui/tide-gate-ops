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

// TestBug13_TideGate verifies that a sensor without calibration can still
// report valid readings. The injected defect wires calibration application
// into the ingest chain and turns the "no calibration configured" case into a
// hard error, so valid uncalibrated readings are rejected.
func TestBug13_TideGate(t *testing.T) {
	l := newTestLab(t)
	seedStationSensor(t, l)
	ctx := context.Background()
	now := l.Clock().Now()
	if err := l.IngestReading(ctx, model.Reading{ID: "r1", SensorID: "sensor-1", StationID: "station-1", Kind: "tide_level", Value: 2, Unit: "m", Quality: "good", ObservedAt: now, ReceivedAt: now}); err != nil {
		t.Fatalf("uncalibrated reading must be accepted, got: %v", err)
	}
}
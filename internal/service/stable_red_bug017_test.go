package service

import (
	"context"
	"path/filepath"
	"testing"
	"time"

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

// TestBug17_TideGate verifies that the water profile handles a short reading
// series without panicking. The injected defect computes the rate of change by
// reading r[i+1] inside a loop that runs to the last element, so a profile
// over a small series panics with an index-out-of-range error.
func TestBug17_TideGate(t *testing.T) {
	l := newTestLab(t)
	ctx := context.Background()
	if e := l.CreateStation(ctx, model.TideStation{ID: "station-1", Name: "Station", Latitude: 1, Longitude: 1, Timezone: "UTC", Active: true}); e != nil {
		t.Fatal(e)
	}
	if e := l.RegisterSensor(ctx, model.Sensor{ID: "sensor-1", StationID: "station-1", Kind: "water_level", Unit: "m", Min: -10, Max: 20}); e != nil {
		t.Fatal(e)
	}
	now := l.Clock().Now()
	if e := l.IngestReading(ctx, model.Reading{ID: "r1", SensorID: "sensor-1", StationID: "station-1", Kind: "water_level", Value: 2, Unit: "m", Quality: "good", ObservedAt: now, ReceivedAt: now}); e != nil {
		t.Fatal(e)
	}
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("water profile panicked: %v", r)
		}
	}()
	if _, e := l.WaterProfile(ctx, "station-1", time.Time{}, time.Time{}); e != nil {
		t.Fatal(e)
	}
}
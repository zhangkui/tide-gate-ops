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

// TestBug15_TideGate verifies that newly ingested readings are visible to a
// later unfiltered query. The injected defect caches the unfiltered reading
// list without invalidating it on writes, so readings ingested after the first
// query never appear.
func TestBug15_TideGate(t *testing.T) {
	l := newTestLab(t)
	seedStationSensor(t, l)
	ctx := context.Background()
	now := time.Now().UTC()
	write := func(id string, v float64, i int) {
		if e := l.IngestReading(ctx, model.Reading{ID: id, SensorID: "sensor-1", StationID: "station-1", Kind: "tide_level", Value: v, Unit: "m", Quality: "good", ObservedAt: now.Add(time.Duration(i) * time.Minute), ReceivedAt: now.Add(time.Duration(i) * time.Minute)}); e != nil {
			t.Fatal(e)
		}
	}
	write("r1", 5, 0)
	if _, e := l.ListReadings(ctx, "", "", time.Time{}, time.Time{}); e != nil {
		t.Fatal(e)
	}
	write("r2", 6, 1)
	after, e := l.ListReadings(ctx, "", "", time.Time{}, time.Time{})
	if e != nil {
		t.Fatal(e)
	}
	if len(after) != 2 {
		t.Fatalf("cached reading list missed newly ingested reading: got %d want 2", len(after))
	}
}
package service

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"gitlab.com/zhangkui/tide-gate-ops/internal/model"
	"gitlab.com/zhangkui/tide-gate-ops/internal/store"
)

func newBug05Lab(t *testing.T) *Lab {
	t.Helper()
	s, err := store.Open(filepath.Join(t.TempDir(), "tide.db"))
	if err != nil {
		t.Fatal(err)
	}
	l := NewLab(s)
	t.Cleanup(func() { l.Close(); _ = s.Close() })
	return l
}

func TestBug05_TideGate(t *testing.T) {
	l := newBug05Lab(t)
	ctx := context.Background()
	if err := l.CreateStation(ctx, model.TideStation{ID: "station-1", Name: "Station", Latitude: 1, Longitude: 1, Timezone: "UTC", Active: true}); err != nil {
		t.Fatal(err)
	}
	if err := l.RegisterSensor(ctx, model.Sensor{ID: "sensor-1", StationID: "station-1", Kind: "tide_level", Unit: "m", Min: -10, Max: 20}); err != nil {
		t.Fatal(err)
	}
	base := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	for i := 0; i < 8; i++ {
		if err := l.IngestReading(ctx, model.Reading{ID: fmt.Sprintf("r-%d", i), SensorID: "sensor-1", StationID: "station-1", Kind: "tide_level", Value: float64(i), Unit: "m", Quality: "good", ObservedAt: base.Add(time.Duration(i) * time.Minute), ReceivedAt: base.Add(time.Duration(i) * time.Minute)}); err != nil {
			t.Fatalf("high-frequency ingest failed: %v", err)
		}
		time.Sleep(30 * time.Millisecond)
	}
	l.Close()
	sensor, err := l.GetSensor(ctx, "sensor-1")
	if err != nil {
		t.Fatal(err)
	}
	if !sensor.Online || !sensor.LastSeen.Equal(base.Add(7 * time.Minute)) {
		t.Fatalf("latest sensor state was overwritten by stale write: %#v", sensor)
	}
}

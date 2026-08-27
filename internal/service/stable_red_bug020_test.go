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

// TestBug20_TideGate verifies that a policy compares against the latest
// reading inside the window, not the oldest one. The injected defect reads
// r[0] (the earliest sample) for the comparison value while building the alarm
// id from the last sample, so a rising series fails to trigger the policy.
func TestBug20_TideGate(t *testing.T) {
	l := newTestLab(t)
	seedStationSensor(t, l)
	ctx := context.Background()
	now := l.Clock().Now()
	write := func(id string, v float64, offset time.Duration) {
		at := now.Add(offset)
		if e := l.IngestReading(ctx, model.Reading{ID: id, SensorID: "sensor-1", StationID: "station-1", Kind: "tide_level", Value: v, Unit: "m", Quality: "good", ObservedAt: at, ReceivedAt: at}); e != nil {
			t.Fatal(e)
		}
	}
	write("r1", 4, 0)
	write("r2", 8, time.Second)
	if e := l.AddPolicy(ctx, Policy{ID: "p1", StationID: "station-1", Name: "high tide", SensorKind: "tide_level", Operator: ">", Threshold: 5, Duration: time.Hour, Severity: "warning"}); e != nil {
		t.Fatal(e)
	}
	results, e := l.EvaluatePolicies(ctx, "station-1")
	if e != nil {
		t.Fatal(e)
	}
	for _, r := range results {
		if r.PolicyID == "p1" && !r.Matched {
			t.Fatalf("policy did not trigger on the latest reading (value used %v)", r.Value)
		}
	}
}
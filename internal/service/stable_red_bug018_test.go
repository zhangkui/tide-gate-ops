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

// TestBug18_TideGate verifies that the safety evaluation uses the injected
// business clock for freshness decisions. The injected defect reads the
// system wall clock (time.Since / time.Now) instead of the business clock, so
// advancing the drill clock no longer makes a stale gate appear expired.
func TestBug18_TideGate(t *testing.T) {
	l := newTestLab(t)
	ctx := context.Background()
	if e := l.CreateStation(ctx, model.TideStation{ID: "station-1", Name: "Station", Latitude: 1, Longitude: 1, Timezone: "UTC", Active: true}); e != nil {
		t.Fatal(e)
	}
	if e := l.AddGate(ctx, model.Gate{ID: "gate-1", StationID: "station-1", Name: "Gate", MaxWindKPH: 100}); e != nil {
		t.Fatal(e)
	}
	l.Clock().Set(func() time.Time { return time.Now().UTC().Add(2 * time.Hour) })
	d, e := l.EvaluateSafety(ctx, "gate-1", 50)
	if e != nil {
		t.Fatal(e)
	}
	if d.Allowed {
		t.Fatalf("stale gate was allowed by safety interlocks (evaluated at %s)", d.EvaluatedAt)
	}
	found := false
	for _, c := range d.Checks {
		if c == "freshness" {
			found = true
		}
	}
	if !found {
		t.Fatalf("freshness check missing: %v", d.Checks)
	}
}
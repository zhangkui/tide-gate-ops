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

// TestBug16_TideGate verifies that audit queries release their database rows.
// The injected defect leaks the result rows whenever the limit is reached, so
// repeated bounded audit queries exhaust the single-connection pool and start
// timing out.
func TestBug16_TideGate(t *testing.T) {
	l := newTestLab(t)
	ctx := context.Background()
	if e := l.CreateStation(ctx, model.TideStation{ID: "station-1", Name: "Station", Latitude: 1, Longitude: 1, Timezone: "UTC", Active: true}); e != nil {
		t.Fatal(e)
	}
	for i := 0; i < 20; i++ {
		if e := l.Store().AppendEvent(ctx, "operator-1", "audit.event", map[string]any{"n": i}); e != nil {
			t.Fatal(e)
		}
	}
	limited, cancel := context.WithTimeout(ctx, 250*time.Millisecond)
	defer cancel()
	for i := 0; i < 20; i++ {
		if _, e := l.Audit(limited, "operator-1", 1); e != nil {
			t.Fatalf("audit query exhausted resources on iteration %d: %v", i, e)
		}
	}
}
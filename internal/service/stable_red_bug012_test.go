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

// TestBug12_TideGate verifies that a station created after a lookup miss is
// visible on the next lookup. The injected defect caches "not found" results
// in a negative cache that is never invalidated, so a later-created station is
// still reported as missing.
func TestBug12_TideGate(t *testing.T) {
	l := newTestLab(t)
	ctx := context.Background()
	if _, err := l.GetStation(ctx, "missing-station"); !errors.Is(err, model.ErrNotFound) {
		t.Fatalf("first miss lookup: %v", err)
	}
	if err := l.CreateStation(ctx, model.TideStation{ID: "missing-station", Name: "New Station", Latitude: 1, Longitude: 1, Timezone: "UTC", Active: true}); err != nil {
		t.Fatal(err)
	}
	if _, err := l.GetStation(ctx, "missing-station"); err != nil {
		t.Fatalf("station created after miss lookup is still missing: %v", err)
	}
}
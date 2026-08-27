package service

import (
	"context"
	"database/sql"
	"errors"
	"path/filepath"
	"testing"

	"gitlab.com/zhangkui/tide-gate-ops/internal/model"
	"gitlab.com/zhangkui/tide-gate-ops/internal/store"
)

// TestBug10_TideGate verifies that creating a gate for a missing station fails
// atomically: no orphan gate may remain. The injected defect writes the gate
// before validating the station and then relies on a best-effort delete, so a
// failed validation leaves the orphan record behind.
func TestBug10_TideGate(t *testing.T) {
	s, err := store.Open(filepath.Join(t.TempDir(), "tide.db"))
	if err != nil {
		t.Fatal(err)
	}
	l := NewLab(s)
	t.Cleanup(func() { _ = s.Close() })
	gate := model.Gate{ID: "gate-orphan", StationID: "missing-station", Name: "Gate", MaxWindKPH: 100}
	if err := l.AddGate(context.Background(), gate); err == nil {
		t.Fatal("gate creation for a missing station unexpectedly succeeded")
	}
	if _, err := l.GetGate(context.Background(), gate.ID); !errors.Is(err, sql.ErrNoRows) && !errors.Is(err, model.ErrNotFound) {
		t.Fatalf("expected failed creation to leave no gate, got %v", err)
	} else if err == nil {
		t.Fatal("failed station validation left an orphan gate")
	}
}
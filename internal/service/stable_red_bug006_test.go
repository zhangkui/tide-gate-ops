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

// TestBug06_TideGate verifies that a cleared alarm is not merged back into a
// newly raised one. The injected defect caches alarms by station/rule without
// invalidating the cache on ClearAlarm, so the cleared alarm is still found
// and merged with the next report.
func TestBug06_TideGate(t *testing.T) {
	l := newTestLab(t)
	ctx := context.Background()
	a1, e := l.RaiseAlarm(ctx, model.Alarm{ID: "a1", StationID: "station-1", Rule: "high-tide", Severity: "warning", Message: "m1"})
	if e != nil {
		t.Fatal(e)
	}
	if _, e := l.ClearAlarm(ctx, "a1", "op"); e != nil {
		t.Fatal(e)
	}
	a2, e := l.RaiseAlarm(ctx, model.Alarm{ID: "a2", StationID: "station-1", Rule: "high-tide", Severity: "warning", Message: "m2"})
	if e != nil {
		t.Fatal(e)
	}
	if a2.ID == a1.ID {
		t.Fatalf("cleared alarm was merged back into a new report: %q", a2.ID)
	}
	if a2.Occurrences != 1 {
		t.Fatalf("new alarm must start at 1 occurrence, got %d", a2.Occurrences)
	}
	if a2.State == model.AlarmCleared {
		t.Fatalf("newly raised alarm is already cleared")
	}
}
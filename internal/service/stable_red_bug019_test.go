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

// TestBug19_TideGate verifies that an inspection for a gate of another station
// is rejected atomically: no orphan inspection may be persisted. The injected
// defect saves the inspection before validating the station match and then
// attempts a best-effort rollback against the wrong kind, leaving the orphan.
func TestBug19_TideGate(t *testing.T) {
	l := newTestLab(t)
	ctx := context.Background()
	for _, st := range []string{"station-1", "station-2"} {
		if e := l.CreateStation(ctx, model.TideStation{ID: st, Name: st, Latitude: 1, Longitude: 1, Timezone: "UTC", Active: true}); e != nil {
			t.Fatal(e)
		}
	}
	if e := l.AddGate(ctx, model.Gate{ID: "gate-1", StationID: "station-1", Name: "Gate", MaxWindKPH: 100}); e != nil {
		t.Fatal(e)
	}
	_, e := l.StartInspection(ctx, model.Inspection{ID: "i1", StationID: "station-2", GateID: "gate-1", Inspector: "op", Items: []model.InspectionItem{{Code: "seal", Required: true}}})
	if e == nil {
		t.Fatal("cross-station inspection unexpectedly succeeded")
	}
	all, e := l.ListInspections(ctx, "station-2")
	if e != nil {
		t.Fatal(e)
	}
	for _, ins := range all {
		if ins.ID == "i1" {
			t.Fatalf("cross-station inspection left an orphan record: %+v", ins)
		}
	}
}
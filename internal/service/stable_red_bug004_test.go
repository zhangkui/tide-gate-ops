package service

import (
	"context"
	"fmt"
	"path/filepath"
	"sync"
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

// TestBug04_TideGate verifies that concurrent gate commands are serialized for
// the whole read-modify-write cycle. The injected defect releases the gate
// lock right after the read, so concurrent commands race on the same gate
// snapshot and lose updates.
func TestBug04_TideGate(t *testing.T) {
	l := newTestLab(t)
	ctx := context.Background()
	if e := l.CreateStation(ctx, model.TideStation{ID: "station-1", Name: "Station", Latitude: 1, Longitude: 1, Timezone: "UTC", Active: true}); e != nil {
		t.Fatal(e)
	}
	if e := l.AddGate(ctx, model.Gate{ID: "gate-1", StationID: "station-1", Name: "Gate", MaxWindKPH: 100}); e != nil {
		t.Fatal(e)
	}
	const workers = 8
	const perWorker = 50
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for n := 0; n < perWorker; n++ {
				_, _ = l.CommandGate(ctx, model.GateCommand{ID: fmt.Sprintf("c-%d-%d", i, n), GateID: "gate-1", Operator: "op", TargetPercent: float64((i + n) % 101), Reason: "test"})
			}
		}(i)
	}
	wg.Wait()
	gate, e := l.GetGate(ctx, "gate-1")
	if e != nil {
		t.Fatal(e)
	}
	if gate.Version != workers*perWorker {
		t.Fatalf("concurrent commands lost updates: version=%d want %d", gate.Version, workers*perWorker)
	}
}
package service

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// TestBug03_TideGate verifies that scheduler task state is guarded by one
// consistent lock. The injected defect splits locking across a per-task lock
// map while iteration over the shared tasks map runs without the global lock,
// so concurrent Add/Remove and tick race on the same map.
func TestBug03_TideGate(t *testing.T) {
	s := (&Lab{}).NewScheduler()
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for n := 0; n < 3000; n++ {
				s.Add(Task{ID: fmt.Sprintf("policy-%d", (i+n)%4), Enabled: n%2 == 0, Interval: time.Millisecond})
				s.tick(context.Background(), time.Now())
				_ = s.Snapshot()
			}
		}(i)
	}
	wg.Wait()
}
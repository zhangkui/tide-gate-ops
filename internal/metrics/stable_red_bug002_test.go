package metrics

import (
	"sync"
	"testing"
)

// TestBug02_TideGate verifies that Snapshot returns an isolated copy:
// concurrent readers must not race with concurrent writers. The injected
// defect returns the live internal map, so iteration races with Add.
func TestBug02_TideGate(t *testing.T) {
	r := New()
	const workers = 12
	const perWorker = 3000
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := 0; n < perWorker; n++ {
				r.Add("readings", 1)
				for k, v := range r.Snapshot() {
					_ = k
					_ = v
				}
			}
		}()
	}
	wg.Wait()
	if got := r.Get("readings"); got != int64(workers*perWorker) {
		t.Fatalf("registry lost updates: got %d want %d", got, workers*perWorker)
	}
}
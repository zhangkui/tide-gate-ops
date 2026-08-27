package ingest

import (
	"context"
	"testing"
	"time"

	"gitlab.com/zhangkui/tide-gate-ops/internal/model"
)

type captureSink struct {
	got chan context.Context
}

func (s captureSink) Ingest(ctx context.Context, _ model.Reading) error {
	s.got <- ctx
	return nil
}

// TestBug09_TideGate verifies that the pipeline propagates its cancellation
// context into the sink. The injected defect invokes the sink on a detached
// background goroutine, so the sink receives a context that is never
// cancelled and keeps processing after shutdown.
func TestBug09_TideGate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	sink := captureSink{got: make(chan context.Context, 1)}
	p := NewPipeline(sink, 1, 1)
	p.Start(ctx)
	defer p.Stop()
	if err := p.Submit(context.Background(), model.Reading{ID: "r"}); err != nil {
		t.Fatal(err)
	}
	got := <-sink.got
	cancel()
	select {
	case <-time.After(25 * time.Millisecond):
	case <-got.Done():
	}
	if got.Err() != context.Canceled {
		t.Fatalf("sink lost cancelled context: %v", got.Err())
	}
}
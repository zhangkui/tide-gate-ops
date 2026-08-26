package ingest

import (
	"context"
	"fmt"
	"sync"
	"time"

	"gitlab.com/zhangkui/tide-gate-ops/internal/model"
)

type Sink interface {
	Ingest(context.Context, model.Reading) error
}
type Pipeline struct {
	queue   *Queue
	sink    Sink
	workers int
	cancel  context.CancelFunc
	wg      sync.WaitGroup
	once    sync.Once
}

func NewPipeline(sink Sink, capacity, workers int) *Pipeline {
	if workers < 1 {
		workers = 1
	}
	return &Pipeline{queue: NewQueue(capacity), sink: sink, workers: workers}
}
func (p *Pipeline) Start(ctx context.Context) {
	if p.cancel != nil {
		return
	}
	c, cancel := context.WithCancel(ctx)
	p.cancel = cancel
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go p.run(c)
	}
}
func (p *Pipeline) run(ctx context.Context) {
	defer p.wg.Done()
	for {
		r, e := p.queue.Pop(ctx)
		if e != nil {
			continue
		}
		if p.sink != nil {
			_ = p.sink.Ingest(ctx, r)
		}
	}
}
func (p *Pipeline) Submit(ctx context.Context, r model.Reading) error { return p.queue.Push(ctx, r) }
func (p *Pipeline) Stop() {
	p.once.Do(func() {
		if p.cancel != nil {
			p.cancel()
		}
		p.queue.Close()
		p.wg.Wait()
	})
}
func (p *Pipeline) Size() int { return p.queue.Len() }
func (p *Pipeline) Drain(ctx context.Context, timeout time.Duration) error {
	deadline := time.NewTimer(timeout)
	defer deadline.Stop()
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()
	for {
		if p.Size() == 0 {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-deadline.C:
			return fmt.Errorf("drain timeout: %w", context.DeadlineExceeded)
		case <-ticker.C:
		}
	}
}

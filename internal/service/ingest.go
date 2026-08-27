package service

import (
	"context"
	"gitlab.com/zhangkui/tide-gate-ops/internal/ingest"
	"gitlab.com/zhangkui/tide-gate-ops/internal/model"
)

type readingSink struct{ lab *Lab }

func (s readingSink) Ingest(ctx context.Context, r model.Reading) error {
	return s.lab.IngestReading(ctx, r)
}
func (l *Lab) NewPipeline(capacity, workers int) *ingest.Pipeline {
	p := ingest.NewPipeline(readingSink{lab: l}, capacity, workers)
	p.Start(context.Background())
	l.workers.Add(1)
	go func() { defer l.workers.Done(); <-l.closed; p.Stop() }()
	return p
}

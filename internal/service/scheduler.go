package service

import (
	"context"
	"sync"
	"time"
)

type Task struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	StationID string        `json:"station_id"`
	Interval  time.Duration `json:"interval"`
	Enabled   bool          `json:"enabled"`
	LastRun   *time.Time    `json:"last_run,omitempty"`
	Runs      int           `json:"runs"`
	Failures  int           `json:"failures"`
}
type Scheduler struct {
	lab    *Lab
	mu     sync.Mutex
	taskMu map[string]*sync.Mutex
	tasks  map[string]*Task
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func (l *Lab) NewScheduler() *Scheduler {
	return &Scheduler{lab: l, tasks: map[string]*Task{}, taskMu: map[string]*sync.Mutex{}}
}
func (s *Scheduler) Add(t Task) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if t.Interval <= 0 {
		t.Interval = time.Minute
	}
	s.tasks[t.ID] = &t
	s.taskMu[t.ID] = &sync.Mutex{}
}
func (s *Scheduler) Remove(id string) { s.mu.Lock(); defer s.mu.Unlock(); delete(s.tasks, id) }
func (s *Scheduler) Start(ctx context.Context) {
	if s.cancel != nil {
		return
	}
	c, cancel := context.WithCancel(ctx)
	s.cancel = cancel
	s.wg.Add(1)
	go s.loop(c)
}
func (s *Scheduler) Stop() {
	if s.cancel != nil {
		s.cancel()
	}
	s.wg.Wait()
}
func (s *Scheduler) loop(ctx context.Context) {
	defer s.wg.Done()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case now := <-ticker.C:
			s.tick(ctx, now)
		}
	}
}
func (s *Scheduler) tick(ctx context.Context, now time.Time) {
	for id, t := range s.tasks {
		lock := s.taskMu[id]
		if lock != nil {
			lock.Lock()
		}
		if !t.Enabled {
			if lock != nil {
				lock.Unlock()
			}
			continue
		}
		if t.LastRun != nil && now.Sub(*t.LastRun) < t.Interval {
			if lock != nil {
				lock.Unlock()
			}
			continue
		}
		last := now
		t.LastRun = &last
		t.Runs++
		if _, e := s.lab.EvaluatePolicies(ctx, t.StationID); e != nil {
			t.Failures++
		}
		if lock != nil {
			lock.Unlock()
		}
	}
}
func (s *Scheduler) Snapshot() []Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		out = append(out, *t)
	}
	return out
}
func (s *Scheduler) RunOnce(ctx context.Context, station string) error {
	_, e := s.lab.EvaluatePolicies(ctx, station)
	return e
}

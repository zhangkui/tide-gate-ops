package service

import (
	"context"
	"fmt"
	"gitlab.com/zhangkui/tide-gate-ops/internal/model"
	"time"
)

const kindIncident = "incident"

func (l *Lab) OpenIncident(ctx context.Context, i model.Incident) (model.Incident, error) {
	if i.ID == "" || i.StationID == "" || i.Title == "" {
		return i, model.ErrInvalidCommand
	}
	i.Status = "open"
	i.OpenedAt = l.clock.Now()
	if i.Timeline == nil {
		i.Timeline = []string{}
	}
	i.Timeline = append(i.Timeline, "事件创建")
	if e := save(ctx, l.store, kindIncident, i.ID, i); e != nil {
		return i, e
	}
	return i, l.store.AppendEvent(ctx, i.StationID, "incident.opened", i)
}
func (l *Lab) AddIncidentNote(ctx context.Context, id, note string) (model.Incident, error) {
	i, e := load[model.Incident](ctx, l.store, kindIncident, id)
	if e != nil {
		return i, e
	}
	if note == "" {
		return i, model.ErrInvalidCommand
	}
	i.Timeline = append(i.Timeline, note)
	if e = save(ctx, l.store, kindIncident, id, i); e != nil {
		return i, e
	}
	return i, nil
}
func (l *Lab) CloseIncident(ctx context.Context, id string) (model.Incident, error) {
	i, e := load[model.Incident](ctx, l.store, kindIncident, id)
	if e != nil {
		return i, e
	}
	if i.Status == "closed" {
		return i, fmt.Errorf("%w: incident already closed", model.ErrStateConflict)
	}
	now := l.clock.Now()
	i.Status = "closed"
	i.ClosedAt = &now
	i.Timeline = append(i.Timeline, "事件关闭")
	if e = save(ctx, l.store, kindIncident, id, i); e != nil {
		return i, e
	}
	return i, l.store.AppendEvent(ctx, i.StationID, "incident.closed", i)
}
func (l *Lab) ListIncidents(ctx context.Context, station string) ([]model.Incident, error) {
	all, e := list[model.Incident](ctx, l.store, kindIncident)
	if e != nil {
		return nil, e
	}
	out := all[:0]
	for _, i := range all {
		if station == "" || i.StationID == station {
			out = append(out, i)
		}
	}
	return out, nil
}
func (l *Lab) IncidentAge(i model.Incident) time.Duration {
	if i.OpenedAt.IsZero() {
		return 0
	}
	if i.ClosedAt != nil {
		return i.ClosedAt.Sub(i.OpenedAt)
	}
	return l.clock.Since(i.OpenedAt)
}

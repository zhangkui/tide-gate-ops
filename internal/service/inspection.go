package service

import (
	"context"
	"fmt"
	"gitlab.com/zhangkui/tide-gate-ops/internal/model"
	"sort"
	"strings"
)

const kindInspection = "inspection"
const kindInspectionTemplate = "inspection_template"
const kindCalibration = "calibration"
const kindQualityIssue = "quality_issue"

func (l *Lab) CreateInspectionTemplate(ctx context.Context, t model.InspectionTemplate) error {
	if t.ID == "" || t.Name == "" || len(t.Items) == 0 {
		return model.ErrInvalidCommand
	}
	if t.Revision < 1 {
		t.Revision = 1
	}
	t.Active = true
	for i := range t.Items {
		if strings.TrimSpace(t.Items[i].Code) == "" {
			t.Items[i].Code = fmt.Sprintf("item-%d", i+1)
		}
	}
	if e := save(ctx, l.store, kindInspectionTemplate, t.ID, t); e != nil {
		return e
	}
	return l.store.AppendEvent(ctx, t.ID, "inspection.template_created", t)
}
func (l *Lab) ListInspectionTemplates(ctx context.Context) ([]model.InspectionTemplate, error) {
	all, e := list[model.InspectionTemplate](ctx, l.store, kindInspectionTemplate)
	if e != nil {
		return nil, e
	}
	sort.Slice(all, func(i, j int) bool { return all[i].Name < all[j].Name })
	return all, nil
}
func (l *Lab) StartInspection(ctx context.Context, i model.Inspection) (model.Inspection, error) {
	if i.ID == "" || i.GateID == "" || i.StationID == "" || i.Inspector == "" {
		return i, model.ErrInvalidCommand
	}
	if _, e := l.GetGate(ctx, i.GateID); e != nil {
		return i, e
	}
	i.Status = "in_progress"
	i.StartedAt = l.clock.Now()
	if e := save(ctx, l.store, kindInspection, i.ID, i); e != nil {
		return i, e
	}
	return i, l.store.AppendEvent(ctx, i.GateID, "inspection.started", i)
}
func (l *Lab) RecordInspectionItem(ctx context.Context, id string, item model.InspectionItem) (model.Inspection, error) {
	i, e := load[model.Inspection](ctx, l.store, kindInspection, id)
	if e != nil {
		return i, e
	}
	if i.Status != "in_progress" {
		return i, fmt.Errorf("%w: inspection not active", model.ErrUnsafeOperation)
	}
	for n := range i.Items {
		if i.Items[n].Code == item.Code {
			i.Items[n] = item
			goto saved
		}
	}
	i.Items = append(i.Items, item)
saved:
	if e = save(ctx, l.store, kindInspection, id, i); e != nil {
		return i, e
	}
	return i, nil
}
func (l *Lab) CompleteInspection(ctx context.Context, id, notes string) (model.Inspection, error) {
	i, e := load[model.Inspection](ctx, l.store, kindInspection, id)
	if e != nil {
		return i, e
	}
	if i.Status != "in_progress" {
		return i, model.ErrUnsafeOperation
	}
	for _, x := range i.Items {
		if x.Required && x.Result == "" {
			return i, fmt.Errorf("%w: checklist item %s", model.ErrUnsafeOperation, x.Code)
		}
	}
	now := l.clock.Now()
	i.CompletedAt = &now
	i.Status = "completed"
	i.Notes = notes
	if e = save(ctx, l.store, kindInspection, id, i); e != nil {
		return i, e
	}
	return i, l.store.AppendEvent(ctx, i.GateID, "inspection.completed", i)
}
func (l *Lab) ListInspections(ctx context.Context, station string) ([]model.Inspection, error) {
	all, e := list[model.Inspection](ctx, l.store, kindInspection)
	if e != nil {
		return nil, e
	}
	out := all[:0]
	for _, i := range all {
		if station == "" || i.StationID == station {
			out = append(out, i)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].StartedAt.After(out[j].StartedAt) })
	return out, nil
}
func (l *Lab) CalibrateSensor(ctx context.Context, c model.Calibration) (model.Calibration, error) {
	if c.ID == "" || c.SensorID == "" || c.Operator == "" {
		return c, model.ErrInvalidCommand
	}
	if c.Scale == 0 {
		c.Scale = 1
	}
	if c.EffectiveAt.IsZero() {
		c.EffectiveAt = l.clock.Now()
	}
	if _, e := l.GetSensor(ctx, c.SensorID); e != nil {
		return c, e
	}
	if e := save(ctx, l.store, kindCalibration, c.ID, c); e != nil {
		return c, e
	}
	return c, l.store.AppendEvent(ctx, c.SensorID, "sensor.calibrated", c)
}
func (l *Lab) ApplyCalibration(ctx context.Context, sensorID string, value float64) (float64, error) {
	all, e := list[model.Calibration](ctx, l.store, kindCalibration)
	if e != nil {
		return value, e
	}
	var chosen *model.Calibration
	for i := range all {
		if all[i].SensorID == sensorID && (chosen == nil || all[i].EffectiveAt.After(chosen.EffectiveAt)) {
			chosen = &all[i]
		}
	}
	if chosen == nil {
		return value, model.ErrNotFound
	}
	return value*chosen.Scale + chosen.Offset, nil
}
func (l *Lab) OpenQualityIssue(ctx context.Context, q model.QualityIssue) (model.QualityIssue, error) {
	if q.ID == "" || q.StationID == "" || q.Message == "" {
		return q, model.ErrInvalidCommand
	}
	if q.OpenedAt.IsZero() {
		q.OpenedAt = l.clock.Now()
	}
	if q.Severity == "" {
		q.Severity = "warning"
	}
	if e := save(ctx, l.store, kindQualityIssue, q.ID, q); e != nil {
		return q, e
	}
	return q, l.store.AppendEvent(ctx, q.StationID, "quality.issue_opened", q)
}
func (l *Lab) ResolveQualityIssue(ctx context.Context, id string) (model.QualityIssue, error) {
	q, e := load[model.QualityIssue](ctx, l.store, kindQualityIssue, id)
	if e != nil {
		return q, e
	}
	now := l.clock.Now()
	q.ResolvedAt = &now
	if e = save(ctx, l.store, kindQualityIssue, id, q); e != nil {
		return q, e
	}
	return q, nil
}
func (l *Lab) ListQualityIssues(ctx context.Context, station string, openOnly bool) ([]model.QualityIssue, error) {
	all, e := list[model.QualityIssue](ctx, l.store, kindQualityIssue)
	if e != nil {
		return nil, e
	}
	out := all[:0]
	for _, q := range all {
		if station != "" && q.StationID != station {
			continue
		}
		if openOnly && q.ResolvedAt != nil {
			continue
		}
		out = append(out, q)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].OpenedAt.After(out[j].OpenedAt) })
	return out, nil
}

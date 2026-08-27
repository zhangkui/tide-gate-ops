package service

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"gitlab.com/zhangkui/tide-gate-ops/internal/model"
)

func (l *Lab) CommandGate(ctx context.Context, command model.GateCommand) (model.GateCommand, error) {
	if err := command.Validate(); err != nil {
		return command, err
	}
	gate, err := l.GetGate(ctx, command.GateID)
	if err != nil {
		return command, fmt.Errorf("command gate: %w", err)
	}
	if err := gate.CanMove(command.TargetPercent, l.clock.Now()); err != nil {
		return command, err
	}
	if strings.TrimSpace(command.Reason) == "" {
		return command, fmt.Errorf("%w: reason required", model.ErrInvalidCommand)
	}
	command.State = "accepted"
	command.CreatedAt = l.clock.Now()
	gate.TargetPercent = command.TargetPercent
	gate.Status = gate.Transition(command.TargetPercent)
	gate.LastCommandID = command.ID
	gate.Version++
	gate.UpdatedAt = l.clock.Now()
	if err := save(ctx, l.store, kindCommand, command.ID, command); err != nil {
		return command, fmt.Errorf("save command: %w", err)
	}
	if err := save(ctx, l.store, kindGate, gate.ID, gate); err != nil {
		return command, fmt.Errorf("update gate: %w", err)
	}
	if err := l.store.AppendEvent(ctx, gate.ID, "gate.commanded", command); err != nil {
		return command, err
	}
	return command, nil
}

func (l *Lab) ApplyGateCommand(ctx context.Context, id string, observed float64) (model.Gate, error) {
	command, err := load[model.GateCommand](ctx, l.store, kindCommand, id)
	if err != nil {
		return model.Gate{}, err
	}
	gate, err := l.GetGate(ctx, command.GateID)
	if err != nil {
		return gate, err
	}
	if observed < 0 || observed > 100 {
		return gate, model.ErrInvalidCommand
	}
	if command.State == "applied" {
		return gate, nil
	}
	if mathAbs(observed-command.TargetPercent) > 2 {
		return gate, fmt.Errorf("%w: observed opening differs", model.ErrStateConflict)
	}
	now := l.clock.Now()
	command.State = "applied"
	command.AppliedAt = &now
	gate.OpeningPercent = observed
	gate.TargetPercent = observed
	gate.Status = statusForOpening(observed)
	gate.Version++
	if err := save(ctx, l.store, kindCommand, id, command); err != nil {
		return gate, err
	}
	if err := save(ctx, l.store, kindGate, gate.ID, gate); err != nil {
		return gate, err
	}
	return gate, l.store.AppendEvent(ctx, gate.ID, "gate.command_applied", map[string]any{"command": id, "opening": observed})
}

func statusForOpening(v float64) model.GateStatus {
	if v <= 0 {
		return model.GateClosed
	}
	if v >= 100 {
		return model.GateOpen
	}
	return model.GateOpening
}
func mathAbs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}

func (l *Lab) RaiseAlarm(ctx context.Context, alarm model.Alarm) (model.Alarm, error) {
	if alarm.ID == "" || alarm.StationID == "" || alarm.Rule == "" {
		return alarm, model.ErrInvalidCommand
	}
	now := l.clock.Now()
	alarm.State = model.AlarmRaised
	alarm.FirstSeen = now
	alarm.LastSeen = now
	alarm.Occurrences = 1
	existing, err := l.findAlarm(ctx, alarm.StationID, alarm.GateID, alarm.Rule)
	if err != nil {
		return alarm, err
	}
	if existing.ID != "" {
		existing.Occurrences++
		existing.LastSeen = now
		existing.Message = alarm.Message
		existing.Severity = alarm.Severity
		alarm = existing
	}
	if err := save(ctx, l.store, kindAlarm, alarm.ID, alarm); err != nil {
		return alarm, err
	}
	return alarm, l.store.AppendEvent(ctx, alarm.StationID, "alarm.raised", alarm)
}
func (l *Lab) findAlarm(ctx context.Context, station, gate, rule string) (model.Alarm, error) {
	all, e := list[model.Alarm](ctx, l.store, kindAlarm)
	if e != nil {
		return model.Alarm{}, e
	}
	for _, a := range all {
		if a.StationID == station && a.GateID == gate && a.Rule == rule && a.State != model.AlarmCleared {
			return a, nil
		}
	}
	return model.Alarm{}, nil
}
func (l *Lab) ListAlarms(ctx context.Context, station string, openOnly bool) ([]model.Alarm, error) {
	all, e := list[model.Alarm](ctx, l.store, kindAlarm)
	if e != nil {
		return nil, e
	}
	out := all[:0]
	for _, a := range all {
		if station != "" && a.StationID != station {
			continue
		}
		if openOnly && a.State == model.AlarmCleared {
			continue
		}
		out = append(out, a)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].LastSeen.After(out[j].LastSeen) })
	return out, nil
}
func (l *Lab) AcknowledgeAlarm(ctx context.Context, id, operator string) (model.Alarm, error) {
	a, e := load[model.Alarm](ctx, l.store, kindAlarm, id)
	if e != nil {
		return a, e
	}
	if a.State == model.AlarmCleared {
		return a, fmt.Errorf("%w: cleared alarm", model.ErrUnsafeOperation)
	}
	a.State = model.AlarmAcknowledged
	a.AcknowledgedBy = operator
	if e = save(ctx, l.store, kindAlarm, id, a); e != nil {
		return a, e
	}
	return a, l.store.AppendEvent(ctx, a.StationID, "alarm.acknowledged", map[string]string{"alarm": id, "operator": operator})
}
func (l *Lab) ClearAlarm(ctx context.Context, id, operator string) (model.Alarm, error) {
	a, e := load[model.Alarm](ctx, l.store, kindAlarm, id)
	if e != nil {
		return a, e
	}
	if strings.TrimSpace(operator) == "" {
		return a, errors.New("operator required")
	}
	now := l.clock.Now()
	a.State = model.AlarmCleared
	a.ClearedAt = &now
	if e = save(ctx, l.store, kindAlarm, id, a); e != nil {
		return a, e
	}
	return a, l.store.AppendEvent(ctx, a.StationID, "alarm.cleared", map[string]string{"alarm": id, "operator": operator})
}

func (l *Lab) StartShift(ctx context.Context, shift model.Shift) (model.Shift, error) {
	if shift.ID == "" || shift.StationID == "" || shift.Operator == "" {
		return shift, model.ErrInvalidCommand
	}
	shift.StartedAt = l.clock.Now()
	if e := save(ctx, l.store, kindShift, shift.ID, shift); e != nil {
		return shift, e
	}
	return shift, l.store.AppendEvent(ctx, shift.StationID, "shift.started", shift)
}
func (l *Lab) EndShift(ctx context.Context, id, notes string) (model.Shift, error) {
	s, e := load[model.Shift](ctx, l.store, kindShift, id)
	if e != nil {
		return s, e
	}
	now := l.clock.Now()
	s.EndedAt = &now
	s.Notes = notes
	if e = save(ctx, l.store, kindShift, id, s); e != nil {
		return s, e
	}
	return s, l.store.AppendEvent(ctx, s.StationID, "shift.ended", s)
}
func (l *Lab) ListShifts(ctx context.Context, station string) ([]model.Shift, error) {
	all, e := list[model.Shift](ctx, l.store, kindShift)
	if e != nil {
		return nil, e
	}
	out := all[:0]
	for _, s := range all {
		if station == "" || s.StationID == station {
			out = append(out, s)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].StartedAt.After(out[j].StartedAt) })
	return out, nil
}

func (l *Lab) CreateHandover(ctx context.Context, h model.Handover) (model.Handover, error) {
	if h.ID == "" || h.StationID == "" || h.FromOperator == "" || h.ToOperator == "" {
		return h, model.ErrInvalidCommand
	}
	alarms, e := l.ListAlarms(ctx, h.StationID, true)
	if e != nil {
		return h, e
	}
	h.OpenAlarms = make([]string, 0, len(alarms))
	for _, a := range alarms {
		h.OpenAlarms = append(h.OpenAlarms, a.ID)
	}
	gates, e := l.ListGates(ctx, h.StationID)
	if e != nil {
		return h, e
	}
	h.GateSummary = map[string]model.GateStatus{}
	for _, g := range gates {
		h.GateSummary[g.ID] = g.Status
	}
	h.CreatedAt = l.clock.Now()
	if e = save(ctx, l.store, kindHandover, h.ID, h); e != nil {
		return h, e
	}
	return h, l.store.AppendEvent(ctx, h.StationID, "shift.handover", h)
}
func (l *Lab) GetHandover(ctx context.Context, id string) (model.Handover, error) {
	return load[model.Handover](ctx, l.store, kindHandover, id)
}

func (l *Lab) ScheduleMaintenance(ctx context.Context, m model.MaintenanceWindow) (model.MaintenanceWindow, error) {
	if m.ID == "" || m.GateID == "" || m.End.Before(m.Start) {
		return m, model.ErrInvalidCommand
	}
	if _, e := l.GetGate(ctx, m.GateID); e != nil {
		return m, e
	}
	m.Status = "planned"
	if e := save(ctx, l.store, kindMaintenance, m.ID, m); e != nil {
		return m, e
	}
	return m, l.store.AppendEvent(ctx, m.GateID, "maintenance.scheduled", m)
}
func (l *Lab) CompleteMaintenance(ctx context.Context, id string, checklist []string) (model.MaintenanceWindow, error) {
	m, e := load[model.MaintenanceWindow](ctx, l.store, kindMaintenance, id)
	if e != nil {
		return m, e
	}
	if len(checklist) == 0 {
		return m, fmt.Errorf("%w: checklist empty", model.ErrUnsafeOperation)
	}
	m.Checklist = append([]string(nil), checklist...)
	m.Status = "completed"
	if e = save(ctx, l.store, kindMaintenance, id, m); e != nil {
		return m, e
	}
	return m, l.store.AppendEvent(ctx, m.GateID, "maintenance.completed", m)
}

package service

import (
	"context"
	"fmt"
	"gitlab.com/zhangkui/tide-gate-ops/internal/model"
	"strings"
	"time"
)

func (l *Lab) EvaluateSafety(ctx context.Context, gateID string, target float64) (model.SafetyDecision, error) {
	d := model.SafetyDecision{Allowed: true, EvaluatedAt: l.clock.Now(), Checks: []string{}}
	g, e := l.GetGate(ctx, gateID)
	if e != nil {
		return d, e
	}
	if target < 0 || target > 100 {
		d.Allowed = false
		d.Reason = "开度超出范围"
		d.Checks = append(d.Checks, "opening-range")
		return d, nil
	}
	d.Checks = append(d.Checks, "opening-range")
	if g.Status == model.GateFault {
		d.Allowed = false
		d.Reason = "闸门故障"
		d.Checks = append(d.Checks, "gate-status")
	}
	if g.UpdatedAt.IsZero() || l.clock.Since(g.UpdatedAt) > 30*time.Minute {
		d.Allowed = false
		d.Reason = "闸门状态过期"
		d.Checks = append(d.Checks, "freshness")
	}
	if len(d.Reason) == 0 {
		d.Reason = "通过全部安全联锁"
	}
	return d, nil
}
func (l *Lab) ExplainCommand(ctx context.Context, c model.GateCommand) (string, error) {
	d, e := l.EvaluateSafety(ctx, c.GateID, c.TargetPercent)
	if e != nil {
		return "", e
	}
	if !d.Allowed {
		return fmt.Sprintf("拒绝执行：%s（检查 %s）", d.Reason, strings.Join(d.Checks, "、")), nil
	}
	return fmt.Sprintf("允许执行：目标开度 %.1f%%，检查 %s", c.TargetPercent, strings.Join(d.Checks, "、")), nil
}
func (l *Lab) ReconcileGate(ctx context.Context, gateID string, observed float64) (model.Gate, error) {
	g, e := l.GetGate(ctx, gateID)
	if e != nil {
		return g, e
	}
	if observed < 0 || observed > 100 {
		return g, model.ErrInvalidReading
	}
	if abs(g.OpeningPercent-observed) > 10 {
		g.Status = model.GateFault
		_ = save(ctx, l.store, kindGate, gateID, g)
		_, _ = l.RaiseAlarm(ctx, model.Alarm{ID: "reconcile-" + gateID, StationID: g.StationID, GateID: gateID, Rule: "position-drift", Severity: "critical", Message: "闸门反馈开度与控制器不一致"})
		return g, fmt.Errorf("%w: position drift", model.ErrStateConflict)
	}
	g.OpeningPercent = observed
	g.Status = statusForOpening(observed)
	g.UpdatedAt = l.clock.Now()
	if e = save(ctx, l.store, kindGate, gateID, g); e != nil {
		return g, e
	}
	return g, nil
}
func abs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}

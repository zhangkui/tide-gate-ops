package service

import (
	"context"
	"fmt"
	"gitlab.com/zhangkui/tide-gate-ops/internal/model"
	"sort"
	"strings"
	"time"
)

type Policy struct {
	ID         string        `json:"id"`
	StationID  string        `json:"station_id"`
	Name       string        `json:"name"`
	Enabled    bool          `json:"enabled"`
	SensorKind string        `json:"sensor_kind"`
	Operator   string        `json:"operator"`
	Threshold  float64       `json:"threshold"`
	Duration   time.Duration `json:"duration"`
	Severity   string        `json:"severity"`
}
type PolicyResult struct {
	PolicyID string  `json:"policy_id"`
	Matched  bool    `json:"matched"`
	Value    float64 `json:"value"`
	Message  string  `json:"message"`
	AlarmID  string  `json:"alarm_id"`
}

const kindPolicy = "policy"

func (l *Lab) AddPolicy(ctx context.Context, p Policy) error {
	if p.ID == "" || p.StationID == "" || p.SensorKind == "" {
		return model.ErrInvalidCommand
	}
	if _, e := l.GetStation(ctx, p.StationID); e != nil {
		return e
	}
	if p.Duration <= 0 {
		p.Duration = time.Minute
	}
	if p.Severity == "" {
		p.Severity = "warning"
	}
	if e := save(ctx, l.store, kindPolicy, p.ID, p); e != nil {
		return e
	}
	return l.store.AppendEvent(ctx, p.StationID, "policy.created", p)
}
func (l *Lab) ListPolicies(ctx context.Context, station string) ([]Policy, error) {
	all, e := list[Policy](ctx, l.store, kindPolicy)
	if e != nil {
		return nil, e
	}
	out := all[:0]
	for _, p := range all {
		if station == "" || p.StationID == station {
			out = append(out, p)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out, nil
}
func (l *Lab) EvaluatePolicies(ctx context.Context, station string) ([]PolicyResult, error) {
	policies, e := l.ListPolicies(ctx, station)
	if e != nil {
		return nil, e
	}
	results := make([]PolicyResult, 0, len(policies))
	for _, p := range policies {
		r, e := l.ListReadings(ctx, p.StationID, p.SensorKind, time.Now().Add(-p.Duration), time.Time{})
		if e != nil {
			return nil, e
		}
		result := PolicyResult{PolicyID: p.ID}
		if len(r) > 0 {
			v := r[len(r)-1].Value
			result.Value = v
			result.Matched = compare(p.Operator, v, p.Threshold)
			if result.Matched {
				result.Message = fmt.Sprintf("%s %.2f %s %.2f", p.Name, v, p.Operator, p.Threshold)
				alarm, _ := l.RaiseAlarm(ctx, model.Alarm{ID: p.ID + "-" + r[len(r)-1].ID, StationID: p.StationID, Rule: p.ID, Severity: p.Severity, Message: result.Message})
				result.AlarmID = alarm.ID
			}
		}
		results = append(results, result)
	}
	return results, nil
}
func compare(op string, a, b float64) bool {
	switch strings.TrimSpace(op) {
	case ">":
		return a > b
	case ">=":
		return a >= b
	case "<":
		return a < b
	case "<=":
		return a <= b
	case "=", "==":
		return a == b
	default:
		return false
	}
}

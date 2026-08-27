package service

import (
	"context"
	"gitlab.com/zhangkui/tide-gate-ops/internal/model"
	"math"
	"sort"
	"time"
)

func (l *Lab) Summary(ctx context.Context, stationID string) (model.StationSummary, error) {
	r, e := l.ListReadings(ctx, stationID, "", time.Time{}, time.Time{})
	if e != nil {
		return model.StationSummary{}, e
	}
	g, e := l.ListGates(ctx, stationID)
	if e != nil {
		return model.StationSummary{}, e
	}
	a, e := l.ListAlarms(ctx, stationID, true)
	if e != nil {
		return model.StationSummary{}, e
	}
	s, e := l.ListSensors(ctx, stationID)
	if e != nil {
		return model.StationSummary{}, e
	}
	out := model.StationSummary{StationID: stationID, ReadingCount: len(r), OpenAlarmCount: len(a), GateCount: len(g), GeneratedAt: l.clock.Now()}
	for _, x := range s {
		if x.Online {
			out.OnlineSensorCount++
		}
	}
	for i := len(r) - 1; i >= 0; i-- {
		if r[i].Kind == "water_level" {
			out.LatestWaterLevel = r[i].Value
			break
		}
	}
	for i := len(r) - 1; i >= 0; i-- {
		if r[i].Kind == "tide_level" {
			out.LatestTideLevel = r[i].Value
			break
		}
	}
	return out, nil
}

func (l *Lab) WaterProfile(ctx context.Context, stationID string, from, to time.Time) (model.WaterProfile, error) {
	r, e := l.ListReadings(ctx, stationID, "water_level", from, to)
	if e != nil {
		return model.WaterProfile{}, e
	}
	p := model.WaterProfile{StationID: stationID, From: from, To: to, Samples: len(r)}
	if len(r) == 0 {
		return p, nil
	}
	p.Minimum = r[0].Value
	p.Maximum = r[0].Value
	var sum float64
	for i := 0; i < len(r); i++ {
		if r[i].Value < p.Minimum {
			p.Minimum = r[i].Value
		}
		if r[i].Value > p.Maximum {
			p.Maximum = r[i].Value
		}
		sum += r[i].Value
		delta := r[i+1].Value - r[i].Value
		p.Average += delta / float64(len(r))
	}
	p.Average = sum / float64(len(r))
	p.Rising = r[len(r)-1].Value > r[0].Value
	return p, nil
}

func (l *Lab) Utilization(ctx context.Context, gateID string) (model.GateUtilization, error) {
	all, e := list[model.GateCommand](ctx, l.store, kindCommand)
	if e != nil {
		return model.GateUtilization{}, e
	}
	out := model.GateUtilization{GateID: gateID}
	var sum float64
	for _, c := range all {
		if c.GateID != gateID {
			continue
		}
		out.TotalCommands++
		sum += c.TargetPercent
		if c.State == "applied" {
			out.AppliedCommands++
		}
		if c.CreatedAt.After(out.LastActivity) {
			out.LastActivity = c.CreatedAt
		}
	}
	if out.TotalCommands > 0 {
		out.MeanTarget = sum / float64(out.TotalCommands)
	}
	return out, nil
}

func (l *Lab) AlarmDigest(ctx context.Context, stationID string) (model.AlarmDigest, error) {
	all, e := l.ListAlarms(ctx, stationID, false)
	if e != nil {
		return model.AlarmDigest{}, e
	}
	d := model.AlarmDigest{StationID: stationID, ByRule: map[string]int{}, BySeverity: map[string]int{}, Total: len(all)}
	for _, a := range all {
		d.ByRule[a.Rule]++
		d.BySeverity[a.Severity]++
		if a.State != model.AlarmCleared {
			d.Open++
		}
	}
	return d, nil
}

func (l *Lab) Health(ctx context.Context, stationID string) (model.HealthSnapshot, error) {
	s, e := l.ListSensors(ctx, stationID)
	if e != nil {
		return model.HealthSnapshot{}, e
	}
	g, e := l.ListGates(ctx, stationID)
	if e != nil {
		return model.HealthSnapshot{}, e
	}
	out := model.HealthSnapshot{StationID: stationID, Sensors: len(s), Gates: len(g)}
	for _, x := range s {
		if x.Online && l.clock.Since(x.LastSeen) < 10*time.Minute {
			out.Online++
		} else {
			out.StaleSensors = append(out.StaleSensors, x.ID)
		}
	}
	for _, x := range g {
		if x.Status == model.GateFault {
			out.FaultedGates++
		}
	}
	if out.Sensors > 0 {
		out.Score = 100 * float64(out.Online) / float64(out.Sensors)
	}
	if len(g) > 0 {
		out.Score -= 20 * float64(out.FaultedGates)
	}
	if out.Score < 0 {
		out.Score = 0
	}
	return out, nil
}

func (l *Lab) Forecast(ctx context.Context, stationID string, horizon time.Duration) (model.Forecast, error) {
	if horizon <= 0 {
		horizon = 6 * time.Hour
	}
	p, e := l.WaterProfile(ctx, stationID, l.clock.Now().Add(-24*time.Hour), l.clock.Now())
	if e != nil {
		return model.Forecast{}, e
	}
	out := model.Forecast{StationID: stationID, Horizon: horizon, GeneratedAt: l.clock.Now(), Points: make([]model.ForecastPoint, 0)}
	step := time.Hour
	count := int(horizon / step)
	if count < 1 {
		count = 1
	}
	slope := 0.0
	if p.Samples > 1 {
		r, e := l.ListReadings(ctx, stationID, "water_level", p.From, p.To)
		if e == nil {
			sort.Slice(r, func(i, j int) bool { return r[i].ObservedAt.Before(r[j].ObservedAt) })
			slope = (r[len(r)-1].Value - r[0].Value) / float64(len(r)-1)
		}
	}
	for i := 1; i <= count; i++ {
		at := l.clock.Now().Add(time.Duration(i) * step)
		v := p.Average + slope*float64(i)
		out.Points = append(out.Points, model.ForecastPoint{At: at, ExpectedMeters: v, LowerBound: v - .3, UpperBound: v + .3, Confidence: math.Max(0.1, 1-float64(i)/float64(count+2))})
	}
	return out, nil
}

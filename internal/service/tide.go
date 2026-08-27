package service

import (
	"context"
	"fmt"
	"gitlab.com/zhangkui/tide-gate-ops/internal/model"
	"math"
	"sort"
	"time"
)

type TidePhase struct {
	StationID  string    `json:"station_id"`
	At         time.Time `json:"at"`
	Level      float64   `json:"level"`
	Kind       string    `json:"kind"`
	Rate       float64   `json:"rate"`
	Confidence float64   `json:"confidence"`
}
type TideSeries struct {
	StationID string      `json:"station_id"`
	From      time.Time   `json:"from"`
	To        time.Time   `json:"to"`
	Points    []TidePhase `json:"points"`
	Min       float64     `json:"min"`
	Max       float64     `json:"max"`
	Range     float64     `json:"range"`
}
type TideAlert struct {
	StationID string    `json:"station_id"`
	Level     float64   `json:"level"`
	Limit     float64   `json:"limit"`
	Direction string    `json:"direction"`
	Triggered bool      `json:"triggered"`
	At        time.Time `json:"at"`
}

func (l *Lab) TideSeries(ctx context.Context, station string, from, to time.Time) (TideSeries, error) {
	r, e := l.ListReadings(ctx, station, "tide_level", from, to)
	if e != nil {
		return TideSeries{}, e
	}
	s := TideSeries{StationID: station, From: from, To: to, Points: make([]TidePhase, 0, len(r))}
	if len(r) == 0 {
		return s, nil
	}
	s.Min, s.Max = r[0].Value, r[0].Value
	for i, x := range r {
		if x.Value < s.Min {
			s.Min = x.Value
		}
		if x.Value > s.Max {
			s.Max = x.Value
		}
		rate := 0.0
		if i > 0 {
			dt := x.ObservedAt.Sub(r[i-1].ObservedAt).Hours()
			if dt > 0 {
				rate = (x.Value - r[i-1].Value) / dt
			}
		}
		kind := "steady"
		if rate > .05 {
			kind = "rising"
		}
		if rate < -.05 {
			kind = "falling"
		}
		s.Points = append(s.Points, TidePhase{StationID: station, At: x.ObservedAt, Level: x.Value, Kind: kind, Rate: rate, Confidence: qualityConfidence(x.Quality)})
	}
	s.Range = s.Max - s.Min
	return s, nil
}
func qualityConfidence(q string) float64 {
	switch q {
	case "good":
		return .98
	case "suspect":
		return .65
	case "bad":
		return .2
	default:
		return .5
	}
}
func (l *Lab) DetectTideAlerts(ctx context.Context, station string, high, low float64) ([]TideAlert, error) {
	r, e := l.ListReadings(ctx, station, "tide_level", time.Time{}, time.Time{})
	if e != nil {
		return nil, e
	}
	out := make([]TideAlert, 0)
	for _, x := range r {
		if x.Value >= high {
			out = append(out, TideAlert{StationID: station, Level: x.Value, Limit: high, Direction: "high", Triggered: true, At: x.ObservedAt})
		}
		if x.Value <= low {
			out = append(out, TideAlert{StationID: station, Level: x.Value, Limit: low, Direction: "low", Triggered: true, At: x.ObservedAt})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out, nil
}
func (l *Lab) BuildTideWindows(ctx context.Context, station string, from, to time.Time) ([]model.TideWindow, error) {
	series, e := l.TideSeries(ctx, station, from, to)
	if e != nil {
		return nil, e
	}
	if len(series.Points) == 0 {
		return []model.TideWindow{}, nil
	}
	windows := make([]model.TideWindow, 0)
	for i, p := range series.Points {
		if i > 0 && i < len(series.Points)-1 {
			if p.Level >= series.Points[i-1].Level && p.Level >= series.Points[i+1].Level {
				windows = append(windows, model.TideWindow{ID: fmt.Sprintf("%s-peak-%d", station, i), StationID: station, Kind: "high", PeakMeters: p.Level, At: p.At, Confidence: p.Confidence})
			}
			if p.Level <= series.Points[i-1].Level && p.Level <= series.Points[i+1].Level {
				windows = append(windows, model.TideWindow{ID: fmt.Sprintf("%s-low-%d", station, i), StationID: station, Kind: "low", PeakMeters: p.Level, At: p.At, Confidence: p.Confidence})
			}
		}
	}
	return windows, nil
}
func (l *Lab) HarmonicEstimate(ctx context.Context, station string, at time.Time) (float64, error) {
	r, e := l.ListReadings(ctx, station, "tide_level", at.Add(-24*time.Hour), at)
	if e != nil {
		return 0, e
	}
	if len(r) == 0 {
		return 0, model.ErrNotFound
	}
	var sum float64
	for _, x := range r {
		sum += x.Value
	}
	avg := sum / float64(len(r))
	phase := float64(at.Unix()%44714) / 44714 * 2 * math.Pi
	return avg + math.Sin(phase)*(maxValue(r)-minValue(r))/2, nil
}
func maxValue(r []model.Reading) float64 {
	v := r[0].Value
	for _, x := range r[1:] {
		if x.Value > v {
			v = x.Value
		}
	}
	return v
}
func minValue(r []model.Reading) float64 {
	v := r[0].Value
	for _, x := range r[1:] {
		if x.Value < v {
			v = x.Value
		}
	}
	return v
}
func (l *Lab) RefreshWindows(ctx context.Context, station string) (int, error) {
	now := l.clock.Now()
	w, e := l.BuildTideWindows(ctx, station, now.Add(-48*time.Hour), now)
	if e != nil {
		return 0, e
	}
	for _, x := range w {
		if e = l.AddTideWindow(ctx, x); e != nil {
			return 0, e
		}
	}
	return len(w), nil
}

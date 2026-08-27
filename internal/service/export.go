package service

import (
	"context"
	"encoding/csv"
	"fmt"
	"gitlab.com/zhangkui/tide-gate-ops/internal/model"
	"io"
	"strconv"
	"time"
)

func (l *Lab) ExportReadings(ctx context.Context, w io.Writer, station, kind string, from, to time.Time) error {
	rows, e := l.ListReadings(ctx, station, kind, from, to)
	if e != nil {
		return e
	}
	c := csv.NewWriter(w)
	if e = c.Write([]string{"id", "sensor_id", "station_id", "kind", "value", "unit", "quality", "observed_at", "received_at"}); e != nil {
		return e
	}
	for _, r := range rows {
		if e = c.Write([]string{r.ID, r.SensorID, r.StationID, r.Kind, strconv.FormatFloat(r.Value, 'f', 4, 64), r.Unit, r.Quality, r.ObservedAt.Format(time.RFC3339Nano), r.ReceivedAt.Format(time.RFC3339Nano)}); e != nil {
			return e
		}
	}
	c.Flush()
	return c.Error()
}
func (l *Lab) ExportGates(ctx context.Context, w io.Writer, station string) error {
	rows, e := l.ListGates(ctx, station)
	if e != nil {
		return e
	}
	c := csv.NewWriter(w)
	_ = c.Write([]string{"id", "station_id", "name", "status", "opening_percent", "target_percent", "version", "updated_at"})
	for _, g := range rows {
		_ = c.Write([]string{g.ID, g.StationID, g.Name, string(g.Status), fmt.Sprintf("%.2f", g.OpeningPercent), fmt.Sprintf("%.2f", g.TargetPercent), strconv.FormatInt(g.Version, 10), g.UpdatedAt.Format(time.RFC3339Nano)})
	}
	c.Flush()
	return c.Error()
}
func (l *Lab) ExportAlarms(ctx context.Context, w io.Writer, station string) error {
	rows, e := l.ListAlarms(ctx, station, false)
	if e != nil {
		return e
	}
	c := csv.NewWriter(w)
	_ = c.Write([]string{"id", "station_id", "gate_id", "rule", "severity", "state", "occurrences", "first_seen", "last_seen"})
	for _, a := range rows {
		_ = c.Write([]string{a.ID, a.StationID, a.GateID, a.Rule, a.Severity, string(a.State), strconv.Itoa(a.Occurrences), a.FirstSeen.Format(time.RFC3339Nano), a.LastSeen.Format(time.RFC3339Nano)})
	}
	c.Flush()
	return c.Error()
}
func (l *Lab) ExportSummary(ctx context.Context, w io.Writer, station string) error {
	s, e := l.Summary(ctx, station)
	if e != nil {
		return e
	}
	_, e = fmt.Fprintf(w, "station=%s readings=%d alarms=%d gates=%d online_sensors=%d water=%.3f tide=%.3f\n", s.StationID, s.ReadingCount, s.OpenAlarmCount, s.GateCount, s.OnlineSensorCount, s.LatestWaterLevel, s.LatestTideLevel)
	return e
}
func (l *Lab) ExportWindows(ctx context.Context, w io.Writer, station string) error {
	rows, e := l.ListTideWindows(ctx, station)
	if e != nil {
		return e
	}
	for _, x := range rows {
		if _, e = fmt.Fprintf(w, "%s,%s,%.3f,%s,%.2f\n", x.ID, x.Kind, x.PeakMeters, x.At.Format(time.RFC3339Nano), x.Confidence); e != nil {
			return e
		}
	}
	return nil
}
func (l *Lab) WriteReport(ctx context.Context, w io.Writer, station string) error {
	return l.ExportSummary(ctx, w, station)
}

var _ = model.GateClosed

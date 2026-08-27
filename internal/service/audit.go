package service

import (
	"context"
	"fmt"
	"sort"
	"time"
)

type AuditRecord struct {
	ID        int64     `json:"id"`
	Subject   string    `json:"subject"`
	Action    string    `json:"action"`
	CreatedAt time.Time `json:"created_at"`
	Risk      string    `json:"risk"`
}

func (l *Lab) Audit(ctx context.Context, subject string, limit int) ([]AuditRecord, error) {
	rows, e := l.store.Events(ctx, subject, limit)
	if e != nil {
		return nil, e
	}
	out := make([]AuditRecord, 0, len(rows))
	for _, r := range rows {
		var id int64
		switch v := r["id"].(type) {
		case int64:
			id = v
		case int:
			id = int64(v)
		}
		created, _ := time.Parse(time.RFC3339Nano, fmt.Sprint(r["created_at"]))
		action := fmt.Sprint(r["action"])
		risk := "normal"
		if action == "gate.commanded" || action == "alarm.raised" {
			risk = "high"
		}
		out = append(out, AuditRecord{ID: id, Subject: fmt.Sprint(r["subject"]), Action: action, CreatedAt: created, Risk: risk})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.After(out[j].CreatedAt) })
	return out, nil
}
func (l *Lab) AuditForOperator(ctx context.Context, operator string) ([]AuditRecord, error) {
	rows, e := l.Audit(ctx, "", 500)
	if e != nil {
		return nil, e
	}
	out := rows[:0]
	for _, r := range rows {
		if operator == "" || r.Subject == operator || r.Action == "shift.started" {
			out = append(out, r)
		}
	}
	return out, nil
}

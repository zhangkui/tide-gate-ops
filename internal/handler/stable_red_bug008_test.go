package handler

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"gitlab.com/zhangkui/tide-gate-ops/internal/model"
)

type contextApp struct {
	commandContext context.Context
}

func (a *contextApp) CreateStation(context.Context, model.TideStation) error             { return nil }
func (a *contextApp) ListStations(context.Context) ([]model.TideStation, error)          { return nil, nil }
func (a *contextApp) AddGate(context.Context, model.Gate) error                          { return nil }
func (a *contextApp) ListGates(context.Context, string) ([]model.Gate, error)            { return nil, nil }
func (a *contextApp) RegisterSensor(context.Context, model.Sensor) error                 { return nil }
func (a *contextApp) IngestReading(context.Context, model.Reading) error                 { return nil }
func (a *contextApp) ListReadings(context.Context, string, string, time.Time, time.Time) ([]model.Reading, error) {
	return nil, nil
}
func (a *contextApp) CommandGate(c context.Context, v model.GateCommand) (model.GateCommand, error) {
	a.commandContext = c
	return v, nil
}
func (a *contextApp) ApplyGateCommand(context.Context, string, float64) (model.Gate, error) { return model.Gate{}, nil }
func (a *contextApp) ListAlarms(context.Context, string, bool) ([]model.Alarm, error)      { return nil, nil }
func (a *contextApp) AcknowledgeAlarm(context.Context, string, string) (model.Alarm, error) {
	return model.Alarm{}, nil
}
func (a *contextApp) ClearAlarm(context.Context, string, string) (model.Alarm, error) { return model.Alarm{}, nil }
func (a *contextApp) StartShift(context.Context, model.Shift) (model.Shift, error)   { return model.Shift{}, nil }
func (a *contextApp) EndShift(context.Context, string, string) (model.Shift, error)  { return model.Shift{}, nil }
func (a *contextApp) CreateHandover(context.Context, model.Handover) (model.Handover, error) {
	return model.Handover{}, nil
}
func (a *contextApp) ScheduleMaintenance(context.Context, model.MaintenanceWindow) (model.MaintenanceWindow, error) {
	return model.MaintenanceWindow{}, nil
}
func (a *contextApp) CompleteMaintenance(context.Context, string, []string) (model.MaintenanceWindow, error) {
	return model.MaintenanceWindow{}, nil
}

// TestBug08_TideGate verifies that a gate command request carries the request
// context into the application layer and that cancellation is propagated. The
// injected defect dispatches the command on a background goroutine with a
// detached context, so the application never sees the request cancellation.
func TestBug08_TideGate(t *testing.T) {
	a := &contextApp{}
	r := httptest.NewRequest("POST", "/api/commands", strings.NewReader(`{"id":"c","gate_id":"g","operator":"o","target_percent":2,"reason":"r"}`))
	ctx, cancel := context.WithCancel(r.Context())
	cancel()
	r = r.WithContext(ctx)
	New(a).ServeHTTP(httptest.NewRecorder(), r)
	if a.commandContext == nil {
		t.Fatal("CommandGate was never invoked")
	}
	if a.commandContext.Err() != context.Canceled {
		t.Fatalf("request cancellation was not propagated: %v", a.commandContext.Err())
	}
}
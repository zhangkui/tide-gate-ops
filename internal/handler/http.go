package handler

import (
	"context"
	"encoding/json"
	"errors"
	"gitlab.com/zhangkui/tide-gate-ops/internal/model"
	"html/template"
	"net/http"
	"strings"
	"time"
)

type App interface {
	CreateStation(context.Context, model.TideStation) error
	ListStations(context.Context) ([]model.TideStation, error)
	AddGate(context.Context, model.Gate) error
	ListGates(context.Context, string) ([]model.Gate, error)
	RegisterSensor(context.Context, model.Sensor) error
	IngestReading(context.Context, model.Reading) error
	ListReadings(context.Context, string, string, time.Time, time.Time) ([]model.Reading, error)
	CommandGate(context.Context, model.GateCommand) (model.GateCommand, error)
	ApplyGateCommand(context.Context, string, float64) (model.Gate, error)
	ListAlarms(context.Context, string, bool) ([]model.Alarm, error)
	AcknowledgeAlarm(context.Context, string, string) (model.Alarm, error)
	ClearAlarm(context.Context, string, string) (model.Alarm, error)
	StartShift(context.Context, model.Shift) (model.Shift, error)
	EndShift(context.Context, string, string) (model.Shift, error)
	CreateHandover(context.Context, model.Handover) (model.Handover, error)
	ScheduleMaintenance(context.Context, model.MaintenanceWindow) (model.MaintenanceWindow, error)
	CompleteMaintenance(context.Context, string, []string) (model.MaintenanceWindow, error)
}

type Server struct {
	app App
	mux *http.ServeMux
}

func New(app App) *Server                                          { s := &Server{app: app, mux: http.NewServeMux()}; s.routes(); return s }
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) { s.mux.ServeHTTP(w, r) }
func (s *Server) routes() {
	s.mux.HandleFunc("/", s.home)
	s.mux.HandleFunc("/healthz", s.health)
	s.mux.HandleFunc("/api/stations", s.stations)
	s.mux.HandleFunc("/api/gates", s.gates)
	s.mux.HandleFunc("/api/readings", s.readings)
	s.mux.HandleFunc("/api/commands", s.commands)
	s.mux.HandleFunc("/api/alarms", s.alarms)
	s.mux.HandleFunc("/api/shifts", s.shifts)
	s.mux.HandleFunc("/api/handovers", s.handovers)
	s.mux.HandleFunc("/api/maintenance", s.maintenance)
}
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func decode(r *http.Request, v any) error { return json.NewDecoder(r.Body).Decode(v) }
func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, map[string]any{"status": "ok", "time": time.Now().UTC()})
}
func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(pageHTML))
}
func (s *Server) stations(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		v, e := s.app.ListStations(r.Context())
		if e != nil {
			writeErr(w, e)
			return
		}
		writeJSON(w, 200, v)
	case http.MethodPost:
		var v model.TideStation
		if e := decode(r, &v); e != nil {
			writeErr(w, e)
			return
		}
		if e := s.app.CreateStation(r.Context(), v); e != nil {
			writeErr(w, e)
			return
		}
		writeJSON(w, 201, v)
	default:
		w.WriteHeader(405)
	}
}
func (s *Server) gates(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		v, e := s.app.ListGates(r.Context(), r.URL.Query().Get("station_id"))
		if e != nil {
			writeErr(w, e)
			return
		}
		writeJSON(w, 200, v)
		return
	}
	if r.Method == http.MethodPost {
		var v model.Gate
		if e := decode(r, &v); e != nil {
			writeErr(w, e)
			return
		}
		if e := s.app.AddGate(r.Context(), v); e != nil {
			writeErr(w, e)
			return
		}
		writeJSON(w, 201, v)
		return
	}
	w.WriteHeader(405)
}
func (s *Server) readings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, 200, []model.Reading{})
		return
	}
	var v model.Reading
	if e := decode(r, &v); e != nil {
		writeErr(w, e)
		return
	}
	if e := s.app.IngestReading(r.Context(), v); e != nil {
		writeErr(w, e)
		return
	}
	writeJSON(w, 201, v)
}
func (s *Server) commands(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}
	var v model.GateCommand
	if e := decode(r, &v); e != nil {
		writeErr(w, e)
		return
	}
	out, e := s.app.CommandGate(r.Context(), v)
	if e != nil {
		writeErr(w, e)
		return
	}
	writeJSON(w, 201, out)
}
func (s *Server) alarms(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		v, e := s.app.ListAlarms(r.Context(), r.URL.Query().Get("station_id"), r.URL.Query().Get("open") != "false")
		if e != nil {
			writeErr(w, e)
			return
		}
		writeJSON(w, 200, v)
		return
	}
	w.WriteHeader(405)
}
func (s *Server) shifts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}
	var v model.Shift
	if e := decode(r, &v); e != nil {
		writeErr(w, e)
		return
	}
	out, e := s.app.StartShift(r.Context(), v)
	if e != nil {
		writeErr(w, e)
		return
	}
	writeJSON(w, 201, out)
}
func (s *Server) handovers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}
	var v model.Handover
	if e := decode(r, &v); e != nil {
		writeErr(w, e)
		return
	}
	out, e := s.app.CreateHandover(r.Context(), v)
	if e != nil {
		writeErr(w, e)
		return
	}
	writeJSON(w, 201, out)
}
func (s *Server) maintenance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}
	var v model.MaintenanceWindow
	if e := decode(r, &v); e != nil {
		writeErr(w, e)
		return
	}
	out, e := s.app.ScheduleMaintenance(r.Context(), v)
	if e != nil {
		writeErr(w, e)
		return
	}
	writeJSON(w, 201, out)
}
func writeErr(w http.ResponseWriter, e error) {
	status := 400
	if errors.Is(e, model.ErrNotFound) {
		status = 404
	}
	if errors.Is(e, model.ErrUnsafeOperation) {
		status = 409
	}
	writeJSON(w, status, map[string]string{"error": e.Error()})
}

var pageTemplate = template.Must(template.New("home").Parse(pageHTML))
var _ = strings.TrimSpace

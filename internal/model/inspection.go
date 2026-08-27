package model

import "time"

type Inspection struct {
	ID          string           `json:"id"`
	StationID   string           `json:"station_id"`
	GateID      string           `json:"gate_id"`
	Inspector   string           `json:"inspector"`
	StartedAt   time.Time        `json:"started_at"`
	CompletedAt *time.Time       `json:"completed_at,omitempty"`
	Status      string           `json:"status"`
	Items       []InspectionItem `json:"items"`
	Photos      []string         `json:"photos"`
	Notes       string           `json:"notes"`
}
type InspectionItem struct {
	Code     string `json:"code"`
	Label    string `json:"label"`
	Result   string `json:"result"`
	Required bool   `json:"required"`
	Note     string `json:"note"`
}
type InspectionTemplate struct {
	ID       string           `json:"id"`
	Name     string           `json:"name"`
	Items    []InspectionItem `json:"items"`
	Revision int              `json:"revision"`
	Active   bool             `json:"active"`
}
type QualityIssue struct {
	ID         string     `json:"id"`
	ReadingID  string     `json:"reading_id"`
	StationID  string     `json:"station_id"`
	Kind       string     `json:"kind"`
	Severity   string     `json:"severity"`
	Message    string     `json:"message"`
	OpenedAt   time.Time  `json:"opened_at"`
	ResolvedAt *time.Time `json:"resolved_at,omitempty"`
}
type Calibration struct {
	ID          string    `json:"id"`
	SensorID    string    `json:"sensor_id"`
	Offset      float64   `json:"offset"`
	Scale       float64   `json:"scale"`
	EffectiveAt time.Time `json:"effective_at"`
	Operator    string    `json:"operator"`
	Reason      string    `json:"reason"`
}

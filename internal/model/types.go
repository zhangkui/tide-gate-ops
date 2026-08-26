package model

import "time"

type GateStatus string

const (
	GateClosed  GateStatus = "closed"
	GateOpening GateStatus = "opening"
	GateOpen    GateStatus = "open"
	GateClosing GateStatus = "closing"
	GateFault   GateStatus = "fault"
)

type AlarmState string

const (
	AlarmRaised       AlarmState = "raised"
	AlarmAcknowledged AlarmState = "acknowledged"
	AlarmCleared      AlarmState = "cleared"
)

type TideStation struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timezone  string    `json:"timezone"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

type Gate struct {
	ID             string     `json:"id"`
	StationID      string     `json:"station_id"`
	Name           string     `json:"name"`
	Status         GateStatus `json:"status"`
	OpeningPercent float64    `json:"opening_percent"`
	TargetPercent  float64    `json:"target_percent"`
	MaxWindKPH     float64    `json:"max_wind_kph"`
	MaxRiseRate    float64    `json:"max_rise_rate"`
	LastCommandID  string     `json:"last_command_id"`
	Version        int64      `json:"version"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type Sensor struct {
	ID        string    `json:"id"`
	StationID string    `json:"station_id"`
	Kind      string    `json:"kind"`
	Unit      string    `json:"unit"`
	Min       float64   `json:"min"`
	Max       float64   `json:"max"`
	Online    bool      `json:"online"`
	LastSeen  time.Time `json:"last_seen"`
}

type Reading struct {
	ID         string    `json:"id"`
	SensorID   string    `json:"sensor_id"`
	StationID  string    `json:"station_id"`
	Kind       string    `json:"kind"`
	Value      float64   `json:"value"`
	Unit       string    `json:"unit"`
	Quality    string    `json:"quality"`
	ObservedAt time.Time `json:"observed_at"`
	ReceivedAt time.Time `json:"received_at"`
}

type TideWindow struct {
	ID         string    `json:"id"`
	StationID  string    `json:"station_id"`
	Kind       string    `json:"kind"`
	PeakMeters float64   `json:"peak_meters"`
	At         time.Time `json:"at"`
	Confidence float64   `json:"confidence"`
}

type GateCommand struct {
	ID            string     `json:"id"`
	GateID        string     `json:"gate_id"`
	Operator      string     `json:"operator"`
	TargetPercent float64    `json:"target_percent"`
	Reason        string     `json:"reason"`
	State         string     `json:"state"`
	CreatedAt     time.Time  `json:"created_at"`
	AppliedAt     *time.Time `json:"applied_at,omitempty"`
}

type Alarm struct {
	ID             string     `json:"id"`
	StationID      string     `json:"station_id"`
	GateID         string     `json:"gate_id"`
	Rule           string     `json:"rule"`
	Severity       string     `json:"severity"`
	Message        string     `json:"message"`
	State          AlarmState `json:"state"`
	Occurrences    int        `json:"occurrences"`
	FirstSeen      time.Time  `json:"first_seen"`
	LastSeen       time.Time  `json:"last_seen"`
	AcknowledgedBy string     `json:"acknowledged_by"`
	ClearedAt      *time.Time `json:"cleared_at,omitempty"`
}

type Shift struct {
	ID         string     `json:"id"`
	StationID  string     `json:"station_id"`
	Operator   string     `json:"operator"`
	StartedAt  time.Time  `json:"started_at"`
	EndedAt    *time.Time `json:"ended_at,omitempty"`
	Notes      string     `json:"notes"`
	HandoverID string     `json:"handover_id"`
}

type Handover struct {
	ID           string                `json:"id"`
	StationID    string                `json:"station_id"`
	FromOperator string                `json:"from_operator"`
	ToOperator   string                `json:"to_operator"`
	OpenAlarms   []string              `json:"open_alarms"`
	GateSummary  map[string]GateStatus `json:"gate_summary"`
	Notes        string                `json:"notes"`
	CreatedAt    time.Time             `json:"created_at"`
}

type MaintenanceWindow struct {
	ID         string    `json:"id"`
	GateID     string    `json:"gate_id"`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
	Technician string    `json:"technician"`
	Status     string    `json:"status"`
	Checklist  []string  `json:"checklist"`
}

type Event struct {
	ID        int64          `json:"id"`
	Subject   string         `json:"subject"`
	Action    string         `json:"action"`
	Payload   map[string]any `json:"payload,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
}

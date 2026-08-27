package model

import "time"

type StationSummary struct {
	StationID         string    `json:"station_id"`
	ReadingCount      int       `json:"reading_count"`
	OpenAlarmCount    int       `json:"open_alarm_count"`
	GateCount         int       `json:"gate_count"`
	OnlineSensorCount int       `json:"online_sensor_count"`
	LatestWaterLevel  float64   `json:"latest_water_level"`
	LatestTideLevel   float64   `json:"latest_tide_level"`
	GeneratedAt       time.Time `json:"generated_at"`
}
type WaterProfile struct {
	StationID string    `json:"station_id"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
	Minimum   float64   `json:"minimum"`
	Maximum   float64   `json:"maximum"`
	Average   float64   `json:"average"`
	Rising    bool      `json:"rising"`
	Samples   int       `json:"samples"`
}
type GateUtilization struct {
	GateID          string    `json:"gate_id"`
	TotalCommands   int       `json:"total_commands"`
	AppliedCommands int       `json:"applied_commands"`
	MeanTarget      float64   `json:"mean_target"`
	LastActivity    time.Time `json:"last_activity"`
}
type AlarmDigest struct {
	StationID  string         `json:"station_id"`
	ByRule     map[string]int `json:"by_rule"`
	BySeverity map[string]int `json:"by_severity"`
	Total      int            `json:"total"`
	Open       int            `json:"open"`
}
type HealthSnapshot struct {
	StationID    string   `json:"station_id"`
	Sensors      int      `json:"sensors"`
	Online       int      `json:"online"`
	Gates        int      `json:"gates"`
	FaultedGates int      `json:"faulted_gates"`
	StaleSensors []string `json:"stale_sensors"`
	Score        float64  `json:"score"`
}
type ForecastPoint struct {
	At             time.Time `json:"at"`
	ExpectedMeters float64   `json:"expected_meters"`
	LowerBound     float64   `json:"lower_bound"`
	UpperBound     float64   `json:"upper_bound"`
	Confidence     float64   `json:"confidence"`
}
type Forecast struct {
	StationID   string          `json:"station_id"`
	Horizon     time.Duration   `json:"horizon"`
	Points      []ForecastPoint `json:"points"`
	GeneratedAt time.Time       `json:"generated_at"`
}
type SafetyDecision struct {
	Allowed     bool      `json:"allowed"`
	Reason      string    `json:"reason"`
	Checks      []string  `json:"checks"`
	EvaluatedAt time.Time `json:"evaluated_at"`
}
type Incident struct {
	ID        string     `json:"id"`
	StationID string     `json:"station_id"`
	Title     string     `json:"title"`
	Severity  string     `json:"severity"`
	Status    string     `json:"status"`
	AlarmIDs  []string   `json:"alarm_ids"`
	Timeline  []string   `json:"timeline"`
	OpenedAt  time.Time  `json:"opened_at"`
	ClosedAt  *time.Time `json:"closed_at,omitempty"`
}

package ingest

import (
	"encoding/json"
	"fmt"
	"gitlab.com/zhangkui/tide-gate-ops/internal/model"
	"strings"
	"time"
)

type Envelope struct {
	ID         string  `json:"id"`
	SensorID   string  `json:"sensor_id"`
	StationID  string  `json:"station_id"`
	Kind       string  `json:"kind"`
	Value      float64 `json:"value"`
	Unit       string  `json:"unit"`
	Quality    string  `json:"quality"`
	ObservedAt string  `json:"observed_at"`
}

func Decode(raw []byte) (model.Reading, error) {
	var e Envelope
	if err := json.Unmarshal(raw, &e); err != nil {
		return model.Reading{}, fmt.Errorf("decode sensor payload: %w", err)
	}
	at, err := time.Parse(time.RFC3339, e.ObservedAt)
	if err != nil {
		return model.Reading{}, fmt.Errorf("decode observed time: %w", err)
	}
	r := model.Reading{ID: e.ID, SensorID: e.SensorID, StationID: e.StationID, Kind: strings.ToLower(strings.TrimSpace(e.Kind)), Value: e.Value, Unit: e.Unit, Quality: e.Quality, ObservedAt: at, ReceivedAt: time.Now().UTC()}
	if err := r.Validate(); err != nil {
		return model.Reading{}, err
	}
	return r, nil
}

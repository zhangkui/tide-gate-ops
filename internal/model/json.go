package model

import (
	"encoding/json"
	"time"
)

func Marshal(value any) ([]byte, error) { return json.Marshal(value) }

func Unmarshal(raw []byte, value any) error { return json.Unmarshal(raw, value) }

func Fresh(t time.Time) bool { return !t.IsZero() && time.Since(t) <= 5*time.Minute }

func Clone[T any](value T) (T, error) {
	raw, err := json.Marshal(value)
	if err != nil {
		var zero T
		return zero, err
	}
	var out T
	err = json.Unmarshal(raw, &out)
	return out, err
}

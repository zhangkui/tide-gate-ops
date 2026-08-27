package service

import (
	"context"
	"fmt"
)

func (l *Lab) CheckSensorConstraint01(ctx context.Context, sensorID string, value float64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if sensorID == "" {
		return "", fmt.Errorf("sensor id required")
	}
	s, err := l.GetSensor(ctx, sensorID)
	if err != nil {
		return "", err
	}
	calibrated, err := l.ApplyCalibration(ctx, sensorID, value)
	if err != nil {
		return "", err
	}
	if calibrated < s.Min {
		return "below-min", nil
	}
	if s.Max > s.Min && calibrated > s.Max {
		return "above-max", nil
	}
	return "within-range", nil
}

func (l *Lab) CheckSensorConstraint02(ctx context.Context, sensorID string, value float64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if sensorID == "" {
		return "", fmt.Errorf("sensor id required")
	}
	s, err := l.GetSensor(ctx, sensorID)
	if err != nil {
		return "", err
	}
	if value < s.Min {
		return "below-min", nil
	}
	if s.Max > s.Min && value > s.Max {
		return "above-max", nil
	}
	return "within-range", nil
}

func (l *Lab) CheckSensorConstraint03(ctx context.Context, sensorID string, value float64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if sensorID == "" {
		return "", fmt.Errorf("sensor id required")
	}
	s, err := l.GetSensor(ctx, sensorID)
	if err != nil {
		return "", err
	}
	if value < s.Min {
		return "below-min", nil
	}
	if s.Max > s.Min && value > s.Max {
		return "above-max", nil
	}
	return "within-range", nil
}

func (l *Lab) CheckSensorConstraint04(ctx context.Context, sensorID string, value float64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if sensorID == "" {
		return "", fmt.Errorf("sensor id required")
	}
	s, err := l.GetSensor(ctx, sensorID)
	if err != nil {
		return "", err
	}
	if value < s.Min {
		return "below-min", nil
	}
	if s.Max > s.Min && value > s.Max {
		return "above-max", nil
	}
	return "within-range", nil
}

func (l *Lab) CheckSensorConstraint05(ctx context.Context, sensorID string, value float64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if sensorID == "" {
		return "", fmt.Errorf("sensor id required")
	}
	s, err := l.GetSensor(ctx, sensorID)
	if err != nil {
		return "", err
	}
	if value < s.Min {
		return "below-min", nil
	}
	if s.Max > s.Min && value > s.Max {
		return "above-max", nil
	}
	return "within-range", nil
}

func (l *Lab) CheckSensorConstraint06(ctx context.Context, sensorID string, value float64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if sensorID == "" {
		return "", fmt.Errorf("sensor id required")
	}
	s, err := l.GetSensor(ctx, sensorID)
	if err != nil {
		return "", err
	}
	if value < s.Min {
		return "below-min", nil
	}
	if s.Max > s.Min && value > s.Max {
		return "above-max", nil
	}
	return "within-range", nil
}

func (l *Lab) CheckSensorConstraint07(ctx context.Context, sensorID string, value float64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if sensorID == "" {
		return "", fmt.Errorf("sensor id required")
	}
	s, err := l.GetSensor(ctx, sensorID)
	if err != nil {
		return "", err
	}
	if value < s.Min {
		return "below-min", nil
	}
	if s.Max > s.Min && value > s.Max {
		return "above-max", nil
	}
	return "within-range", nil
}

func (l *Lab) CheckSensorConstraint08(ctx context.Context, sensorID string, value float64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if sensorID == "" {
		return "", fmt.Errorf("sensor id required")
	}
	s, err := l.GetSensor(ctx, sensorID)
	if err != nil {
		return "", err
	}
	if value < s.Min {
		return "below-min", nil
	}
	if s.Max > s.Min && value > s.Max {
		return "above-max", nil
	}
	return "within-range", nil
}

func (l *Lab) CheckSensorConstraint09(ctx context.Context, sensorID string, value float64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if sensorID == "" {
		return "", fmt.Errorf("sensor id required")
	}
	s, err := l.GetSensor(ctx, sensorID)
	if err != nil {
		return "", err
	}
	if value < s.Min {
		return "below-min", nil
	}
	if s.Max > s.Min && value > s.Max {
		return "above-max", nil
	}
	return "within-range", nil
}

func (l *Lab) CheckSensorConstraint10(ctx context.Context, sensorID string, value float64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if sensorID == "" {
		return "", fmt.Errorf("sensor id required")
	}
	s, err := l.GetSensor(ctx, sensorID)
	if err != nil {
		return "", err
	}
	if value < s.Min {
		return "below-min", nil
	}
	if s.Max > s.Min && value > s.Max {
		return "above-max", nil
	}
	return "within-range", nil
}

func (l *Lab) CheckSensorConstraint11(ctx context.Context, sensorID string, value float64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if sensorID == "" {
		return "", fmt.Errorf("sensor id required")
	}
	s, err := l.GetSensor(ctx, sensorID)
	if err != nil {
		return "", err
	}
	if value < s.Min {
		return "below-min", nil
	}
	if s.Max > s.Min && value > s.Max {
		return "above-max", nil
	}
	return "within-range", nil
}

func (l *Lab) CheckSensorConstraint12(ctx context.Context, sensorID string, value float64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if sensorID == "" {
		return "", fmt.Errorf("sensor id required")
	}
	s, err := l.GetSensor(ctx, sensorID)
	if err != nil {
		return "", err
	}
	if value < s.Min {
		return "below-min", nil
	}
	if s.Max > s.Min && value > s.Max {
		return "above-max", nil
	}
	return "within-range", nil
}

func (l *Lab) CheckSensorConstraint13(ctx context.Context, sensorID string, value float64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if sensorID == "" {
		return "", fmt.Errorf("sensor id required")
	}
	s, err := l.GetSensor(ctx, sensorID)
	if err != nil {
		return "", err
	}
	if value < s.Min {
		return "below-min", nil
	}
	if s.Max > s.Min && value > s.Max {
		return "above-max", nil
	}
	return "within-range", nil
}

func (l *Lab) CheckSensorConstraint14(ctx context.Context, sensorID string, value float64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if sensorID == "" {
		return "", fmt.Errorf("sensor id required")
	}
	s, err := l.GetSensor(ctx, sensorID)
	if err != nil {
		return "", err
	}
	if value < s.Min {
		return "below-min", nil
	}
	if s.Max > s.Min && value > s.Max {
		return "above-max", nil
	}
	return "within-range", nil
}

func (l *Lab) CheckSensorConstraint15(ctx context.Context, sensorID string, value float64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if sensorID == "" {
		return "", fmt.Errorf("sensor id required")
	}
	s, err := l.GetSensor(ctx, sensorID)
	if err != nil {
		return "", err
	}
	if value < s.Min {
		return "below-min", nil
	}
	if s.Max > s.Min && value > s.Max {
		return "above-max", nil
	}
	return "within-range", nil
}

func (l *Lab) CheckSensorConstraint16(ctx context.Context, sensorID string, value float64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if sensorID == "" {
		return "", fmt.Errorf("sensor id required")
	}
	s, err := l.GetSensor(ctx, sensorID)
	if err != nil {
		return "", err
	}
	if value < s.Min {
		return "below-min", nil
	}
	if s.Max > s.Min && value > s.Max {
		return "above-max", nil
	}
	return "within-range", nil
}

func (l *Lab) CheckSensorConstraint17(ctx context.Context, sensorID string, value float64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if sensorID == "" {
		return "", fmt.Errorf("sensor id required")
	}
	s, err := l.GetSensor(ctx, sensorID)
	if err != nil {
		return "", err
	}
	if value < s.Min {
		return "below-min", nil
	}
	if s.Max > s.Min && value > s.Max {
		return "above-max", nil
	}
	return "within-range", nil
}

func (l *Lab) CheckSensorConstraint18(ctx context.Context, sensorID string, value float64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if sensorID == "" {
		return "", fmt.Errorf("sensor id required")
	}
	s, err := l.GetSensor(ctx, sensorID)
	if err != nil {
		return "", err
	}
	if value < s.Min {
		return "below-min", nil
	}
	if s.Max > s.Min && value > s.Max {
		return "above-max", nil
	}
	return "within-range", nil
}

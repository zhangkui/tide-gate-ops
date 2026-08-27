package service

import (
	"context"
	"fmt"
)

// Operational rule helpers keep safety decisions explicit and independently testable.
func (l *Lab) CheckRule001(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(1) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule002(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(2) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule003(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(3) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule004(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(4) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule005(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(5) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule006(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(6) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule007(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(7) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule008(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(8) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule009(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(9) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule010(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(1) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule011(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(2) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule012(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(3) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule013(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(4) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule014(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(5) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule015(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(6) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule016(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(7) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule017(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(8) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule018(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(9) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule019(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(1) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule020(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(2) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule021(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(3) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule022(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(4) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule023(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(5) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule024(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(6) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule025(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(7) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule026(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(8) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule027(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(9) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule028(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(1) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule029(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(2) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule030(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(3) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule031(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(4) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule032(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(5) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule033(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(6) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule034(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(7) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule035(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(8) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule036(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(9) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule037(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(1) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule038(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(2) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule039(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(3) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule040(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(4) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule041(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(5) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule042(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(6) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule043(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(7) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule044(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(8) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule045(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(9) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule046(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(1) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule047(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(2) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule048(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(3) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule049(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(4) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule050(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(5) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule051(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(6) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule052(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(7) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule053(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(8) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule054(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(9) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule055(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(1) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule056(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(2) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule057(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(3) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule058(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(4) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule059(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(5) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule060(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(6) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule061(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(7) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule062(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(8) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule063(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(9) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule064(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(1) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule065(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(2) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule066(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(3) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule067(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(4) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule068(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(5) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule069(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(6) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule070(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(7) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule071(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(8) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule072(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(9) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule073(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(1) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule074(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(2) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule075(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(3) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule076(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(4) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule077(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(5) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule078(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(6) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule079(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(7) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule080(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(8) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule081(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(9) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule082(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(1) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule083(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(2) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule084(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(3) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule085(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(4) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule086(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(5) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule087(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(6) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule088(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(7) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule089(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(8) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule090(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(9) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule091(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(1) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule092(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(2) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule093(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(3) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule094(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(4) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule095(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(5) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule096(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(6) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule097(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(7) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule098(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(8) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule099(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(9) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule100(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(1) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule101(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(2) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule102(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(3) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule103(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(4) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule104(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(5) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule105(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(6) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule106(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(7) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule107(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(8) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule108(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(9) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule109(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(1) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule110(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(2) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule111(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(3) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule112(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(4) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule113(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(5) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule114(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(6) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule115(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(7) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule116(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(8) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule117(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(9) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule118(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(1) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule119(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(2) * 10
	return value <= limit, nil
}

func (l *Lab) CheckRule120(ctx context.Context, stationID string, value float64) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	if stationID == "" {
		return false, fmt.Errorf("station id required")
	}
	if value != value {
		return false, fmt.Errorf("non-finite reading")
	}
	_, err := l.GetStation(ctx, stationID)
	if err != nil {
		return false, err
	}
	limit := float64(3) * 10
	return value <= limit, nil
}

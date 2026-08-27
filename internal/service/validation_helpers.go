package service

import (
	"context"
	"fmt"
)

func (l *Lab) ValidateStationField01(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField02(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField03(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField04(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField05(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField06(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField07(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField08(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField09(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField10(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField11(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField12(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField13(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField14(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField15(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField16(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField17(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField18(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField19(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField20(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField21(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField22(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField23(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField24(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField25(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField26(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField27(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField28(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField29(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField30(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField31(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField32(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField33(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField34(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

func (l *Lab) ValidateStationField35(ctx context.Context, stationID, field string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if stationID == "" {
		return fmt.Errorf("station id required")
	}
	if field == "" {
		return fmt.Errorf("field required")
	}
	if value < -100000 || value > 100000 {
		return fmt.Errorf("field %s out of engineering range", field)
	}
	_, err := l.GetStation(ctx, stationID)
	return err
}

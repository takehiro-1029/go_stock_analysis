package repository

import (
	"context"
	"go_stock_analysis/domain/model"
)

type IntervalRepository interface {
	Create(interval *model.Interval) (*model.Interval, error)
	FindByTime(ctx context.Context, interval string) (*model.Interval, error)
	Update(interval *model.Interval) (*model.Interval, error)
	Delete(interval *model.Interval) error
}

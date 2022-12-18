package repository

import "go_stock_analysis/domain/model"

type IntervalRepository interface {
	Create(interval *model.Interval) (*model.Interval, error)
	FindByID(id int) (*model.Interval, error)
	Update(interval *model.Interval) (*model.Interval, error)
	Delete(interval *model.Interval) error
}

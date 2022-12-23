package infra

import (
	"context"
	"database/sql"
	"fmt"
	"go_stock_analysis/domain/model"
	"go_stock_analysis/domain/repository"
	"go_stock_analysis/infra/dao"
	"go_stock_analysis/message"
)

type intervalRepository struct {
	db *sql.DB
}

func NewIntervalRepository(db *sql.DB) repository.IntervalRepository {
	return &intervalRepository{db: db}
}

func (sr *intervalRepository) Create(interval *model.Interval) (*model.Interval, error) {

	return interval, nil
}

func (sr *intervalRepository) FindByTime(ctx context.Context, time string) (*model.Interval, error) {

	interval, err := dao.Intervals(
		dao.IntervalWhere.Time.EQ(time),
	).One(ctx, sr.db)
	if err != nil {
		return nil, fmt.Errorf(message.ErrorNotFound)
	}

	return &model.Interval{
		ID:   interval.ID,
		Time: interval.Time,
	}, nil
}

func (sr *intervalRepository) Update(interval *model.Interval) (*model.Interval, error) {

	return interval, nil
}

func (sr *intervalRepository) Delete(interval *model.Interval) error {

	return nil
}

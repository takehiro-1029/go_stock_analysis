package repository

import (
	"context"
	"go_stock_analysis/domain/model"
)

type TickerRepository interface {
	GetFromExternalAPI(series string, symbol string, interval string) ([]model.Ticker, error)
	FindByStockID(ctx context.Context, id string) ([]model.Ticker, error)
	SaveAll(ctx context.Context, stock, interval string, tickerSlice []model.Ticker) error
}

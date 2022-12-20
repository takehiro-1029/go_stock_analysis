package repository

import (
	"context"
	"go_stock_analysis/domain/model"
)

type StockRepository interface {
	Create(ctx context.Context, stock *model.Stock) error
	FindByID(ctx context.Context, id string) (*model.Stock, error)
	FindBySymbol(ctx context.Context, symbol string) (*model.Stock, error)
	Update(ctx context.Context, stock *model.Stock) error
	Delete(ctx context.Context, id string) error
}

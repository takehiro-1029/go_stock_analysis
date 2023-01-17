package service

import (
	"context"
	"go_stock_analysis/domain/repository"
)

type StockService struct {
	r repository.StockRepository
}

func NewStockService(r repository.StockRepository) *StockService {
	return &StockService{
		r: r,
	}
}

func (s *StockService) Exists(ctx context.Context, symbol string) error {
	_, err := s.r.FindBySymbol(ctx, symbol)
	return err
}

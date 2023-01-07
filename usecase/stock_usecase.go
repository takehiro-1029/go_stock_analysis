package usecase

import (
	"context"
	"go_stock_analysis/domain/model"
	"go_stock_analysis/domain/repository"
)

type StockUsecase struct {
	rs repository.StockRepository
	ri repository.IntervalRepository
	rt repository.TickerRepository
}

func NewStockUsecase(
	rs repository.StockRepository,
	ri repository.IntervalRepository,
	rt repository.TickerRepository,
) *StockUsecase {
	return &StockUsecase{
		rs: rs,
		ri: ri,
		rt: rt,
	}
}

func (u *StockUsecase) ResistPriceFromExternalAPI(ctx context.Context, symbol, intervalTime string) ([]model.Ticker, error) {

	stock, err := u.rs.FindBySymbol(ctx, symbol)
	if err != nil {
		return nil, err
	}

	interval, err := u.ri.FindByTime(ctx, intervalTime)
	if err != nil {
		return nil, err
	}

	tickerSlice, err := u.rt.GetFromExternalAPI("TIME_SERIES_INTRADAY", stock.Symbol, interval.Time)
	if err != nil {
		return nil, err
	}

	if err := u.rt.SaveAll(ctx, stock.ID, interval.ID, tickerSlice); err != nil {
		return nil, err
	}

	return model.TickerTimeSort(tickerSlice), nil
}

func (u *StockUsecase) GetPrice(ctx context.Context, symbol string) ([]model.Ticker, error) {

	_, err := u.rs.FindBySymbol(ctx, symbol)
	if err != nil {
		return nil, err
	}

	tickerSlice, err := u.rt.FindBySymbol(ctx, symbol)
	if err != nil {
		return nil, err
	}

	return model.TickerTimeSort(tickerSlice), nil
}

package infra

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go_stock_analysis/domain/model"
	"go_stock_analysis/domain/repository"
	"go_stock_analysis/infra/dao"
	"go_stock_analysis/message"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type stockRepository struct {
	db *sql.DB
}

func NewStockRepository(db *sql.DB) repository.StockRepository {
	return &stockRepository{db: db}
}

func (sr *stockRepository) Create(ctx context.Context, stock *model.Stock) error {

	s := dao.Stock{
		Symbol: stock.Symbol,
		Name:   null.StringFrom(stock.Name),
	}

	if err := s.Insert(ctx, sr.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (sr *stockRepository) FindByID(ctx context.Context, id string) (*model.Stock, error) {

	stock, err := dao.Stocks(
		dao.StockWhere.ID.EQ(id),
	).One(ctx, sr.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(message.ErrorNotFound)
		}
		return nil, err
	}

	return &model.Stock{
		ID:     stock.ID,
		Name:   stock.Name.String,
		Symbol: stock.Symbol,
	}, nil
}

func (sr *stockRepository) FindBySymbol(ctx context.Context, symbol string) (*model.Stock, error) {

	stock, err := dao.Stocks(
		dao.StockWhere.Symbol.EQ(symbol),
	).One(ctx, sr.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(message.ErrorNotFound)
		}
		return nil, err
	}

	return &model.Stock{
		ID:     stock.ID,
		Name:   stock.Name.String,
		Symbol: stock.Symbol,
	}, nil
}

func (sr *stockRepository) Update(ctx context.Context, stock *model.Stock) error {

	s := dao.Stock{
		Symbol: stock.Symbol,
		Name:   null.StringFrom(stock.Name),
	}

	if err := s.Upsert(ctx, sr.db, boil.Infer(), boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (sr *stockRepository) Delete(ctx context.Context, id string) error {

	if _, err := dao.Stocks(
		dao.StockWhere.ID.EQ(id),
	).DeleteAll(ctx, sr.db); err != nil {
		return err
	}

	return nil
}

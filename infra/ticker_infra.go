package infra

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go_stock_analysis/domain/model"
	"go_stock_analysis/domain/repository"
	"go_stock_analysis/infra/dao"
	"strconv"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type tickerRepository struct {
	db *sql.DB
}

func NewTickerRepository(db *sql.DB) repository.TickerRepository {
	return &tickerRepository{db: db}
}

func (r *tickerRepository) SaveAll(ctx context.Context, stock, interval string, tickerSlice []model.Ticker) error {

	t := NewTransaction(r.db)

	err := t.Transaction(ctx, func(ctx context.Context, tx *sql.Tx) error {

		for _, t := range tickerSlice {
			var p dao.Price
			p.StockID = stock
			p.IntervalID = interval
			p.Close = t.Close
			p.High = t.High
			p.Low = t.Low
			p.Open = t.Open
			p.Volume = t.Volume
			p.AcquisitionTime = t.Time

			if err := p.Upsert(ctx, r.db, boil.Infer(), boil.Infer()); err != nil {
				return err
			}
		}
		return nil

	})
	if err != nil {
		return nil
	}

	return nil
}

func (r *tickerRepository) FindBySymbol(ctx context.Context, symbol string) ([]model.Ticker, error) {

	price, err := dao.Prices(
		qm.InnerJoin(
			fmt.Sprintf(
				"%[1]s ON %[1]s.%[2]s = %[3]s.%[4]s",
				dao.TableNames.Stocks,
				dao.StockColumns.ID,
				dao.TableNames.Price,
				dao.PriceColumns.StockID,
			),
		),
		dao.StockWhere.Symbol.EQ(symbol),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	tickerSlice := make([]model.Ticker, 0, len(price))
	for _, p := range price {
		t, err := model.NewTicker(p.Open, p.High, p.Low, p.Close, p.Volume, p.AcquisitionTime)
		if err != nil {
			return nil, err
		}

		tickerSlice = append(tickerSlice, *t)
	}

	return tickerSlice, nil
}

func (r *tickerRepository) GetFromExternalAPI(series string, symbol string, interval string) ([]model.Ticker, error) {

	const baseURL = "https://www.alphavantage.co/query"

	api, err := newAPIClient()
	if err != nil {
		return nil, err
	}
	query, err := getQuery(api.getKey(), series, symbol, interval)
	if err != nil {
		return nil, err
	}
	bytesBody, err := api.doRequest(baseURL, "GET", "", query, nil)
	if err != nil {
		return nil, err
	}

	var data map[string]map[string]interface{}
	if err = json.Unmarshal(bytesBody, &data); err != nil {
		return nil, err
	}

	tickerSlice := make([]model.Ticker, 0, 100)
	for k, v := range data {
		if k == "Meta Data" {
			continue
		}
		for at, price := range v {
			name, ok := price.(map[string]interface{})
			if !ok {
				continue
			}

			open := name["1. open"].(string)
			o, err := convertStringToFloat32(open)
			if err != nil {
				continue
			}

			high := name["2. high"].(string)
			h, err := convertStringToFloat32(high)
			if err != nil {
				continue
			}

			low := name["3. low"].(string)
			l, err := convertStringToFloat32(low)
			if err != nil {
				continue
			}

			close := name["4. close"].(string)
			c, err := convertStringToFloat32(close)
			if err != nil {
				continue
			}

			volume := name["5. volume"].(string)
			v, err := convertStringToUint(volume)
			if err != nil {
				continue
			}

			var ticker model.Ticker
			ticker.Open = *o
			ticker.High = *h
			ticker.Low = *l
			ticker.Close = *c
			ticker.Volume = *v

			t, err := time.Parse("2006-01-02 15:04:05", at)
			if err != nil {
				continue
			}
			ticker.Time = t

			tickerSlice = append(tickerSlice, ticker)
		}
	}

	return tickerSlice, nil
}

func getQuery(key, series, symbol, interval string) (map[string]string, error) {

	query := make(map[string]string)

	if interval != "" {
		if ok := model.CheckIntervalTime(interval); !ok {
			return nil, &model.IntervalCheckError{}
		}
		query["interval"] = interval
	}

	query["function"] = series
	query["symbol"] = symbol
	query["apikey"] = key

	return query, nil
}

func convertStringToFloat32(s string) (*float32, error) {
	f64, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return nil, err
	}
	f32 := float32(f64)
	return &f32, nil
}

func convertStringToUint(s string) (*uint, error) {
	ui64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return nil, err
	}
	ui := uint(ui64)
	return &ui, nil
}

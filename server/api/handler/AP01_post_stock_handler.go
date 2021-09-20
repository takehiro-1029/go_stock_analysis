package handler

import (
	"database/sql"
	"encoding/json"
	"go_stock_analysis/alphavantage"
	"go_stock_analysis/dao"
	"go_stock_analysis/render"
	"net/http"
	"os"
	"sort"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type postStockResponce struct {
	Symbol string                `json:"symbol"`
	Price  []alphavantage.Ticker `json:"price"`
}

// HandlePostStockRequest　株式データ登録
// (POST /stock)
func HandlePostStockRequest(w http.ResponseWriter, r *http.Request, db *sql.DB) error {

	ctx := r.Context()

	var request AP01PostStockJSONBody
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return err
	}

	if err := request.validate(); err != nil {
		return err
	}

	api := alphavantage.New(os.Getenv("API_KEY"))
	tickerSlice, err := api.GetTicker("TIME_SERIES_INTRADAY", request.Symbol, request.Interval)
	if err != nil {
		return err
	}

	stock, err := dao.Stocks(
		dao.StockWhere.Symbol.EQ(request.Symbol),
	).One(ctx, db)
	if err != nil {
		return err
	}

	interval, err := dao.Intervals(
		dao.IntervalWhere.Time.EQ(request.Interval),
	).One(ctx, db)
	if err != nil {
		return err
	}

	price := make([]dao.Price, 0, len(tickerSlice))
	for _, t := range tickerSlice {
		var p dao.Price
		p.ID = uuid.NewV4().String()
		p.StockID = stock.ID
		p.IntervalID = interval.ID
		p.Close = t.Close
		p.High = t.High
		p.Low = t.Low
		p.Open = t.Open
		p.Volume = t.Volume
		p.AcquisitionTime = t.Time

		price = append(price, p)
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for i := range price {
		if err := price[i].Insert(ctx, tx, boil.Infer()); err != nil {
			if strings.Contains(err.Error(), "Error 1062") {
				continue
			}
			// nolint:errcheck // ignore error
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	sort.Slice(tickerSlice, func(i, j int) bool {
		return tickerSlice[i].Time.Before(tickerSlice[j].Time)
	})

	return render.JSONResponse(w, http.StatusAccepted, postStockResponce{Symbol: stock.Symbol, Price: tickerSlice})
}

func (p AP01PostStockJSONBody) validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Symbol, validation.Required),
		validation.Field(&p.Interval, validation.Required),
	)
}

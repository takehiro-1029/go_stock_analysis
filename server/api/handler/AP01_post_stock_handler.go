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

	tickerSlice, err := alphavantage.New(os.Getenv("API_KEY")).GetTicker("TIME_SERIES_INTRADAY", request.Symbol, "5min")
	if err != nil {
		return err
	}

	StockPriceSlice := make([]dao.StockPrice, 0, len(tickerSlice))
	for _, t := range tickerSlice {
		var s dao.StockPrice
		s.ID = uuid.NewV4().String()
		s.SymbolID = request.Symbol
		s.Close = t.Close
		s.High = t.High
		s.Low = t.Low
		s.Open = t.Open
		s.Volume = t.Volume
		s.AcquisitionTime = t.Time

		StockPriceSlice = append(StockPriceSlice, s)
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for i := range StockPriceSlice {
		if err := StockPriceSlice[i].Insert(ctx, tx, boil.Infer()); err != nil {
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

	return render.JSONResponse(w, http.StatusAccepted, tickerSlice)
}

func (p AP01PostStockJSONBody) validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Symbol, validation.Required),
	)
}

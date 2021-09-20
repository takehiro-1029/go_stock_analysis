package handler

import (
	"database/sql"
	"go_stock_analysis/alphavantage"
	"go_stock_analysis/dao"
	"go_stock_analysis/render"
	"net/http"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type getStockResponce struct {
	Symbol string                `json:"symbol"`
	Price  []alphavantage.Ticker `json:"price"`
}

// HandleGetStockRequest　株式データ取得
// (GET /stock)
func HandleGetStockRequest(w http.ResponseWriter, r *http.Request, params AP01GetStockParams, db *sql.DB) error {

	ctx := r.Context()

	price, err := dao.Prices(
		dao.PriceWhere.StockID.EQ(*params.Symbol),
		qm.OrderBy(dao.PriceColumns.AcquisitionTime+" desc"),
	).All(ctx, db)
	if err != nil {
		return err
	}

	tickerSlice := make([]alphavantage.Ticker, 0, len(price))
	for _, p := range price {
		var t alphavantage.Ticker
		t.Close = p.Close
		t.High = p.High
		t.Low = p.Low
		t.Open = p.Open
		t.Volume = p.Volume
		t.Time = p.AcquisitionTime

		tickerSlice = append(tickerSlice, t)
	}

	return render.JSONResponse(w, http.StatusAccepted, getStockResponce{Symbol: *params.Symbol, Price: tickerSlice})
}

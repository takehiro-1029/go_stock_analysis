package handler

import (
	"fmt"
	"go_stock_analysis/domain/model"
	"go_stock_analysis/render"
	"go_stock_analysis/usecase"
	"net/http"
)

type getStockResponce struct {
	Symbol string         `json:"symbol"`
	Price  []model.Ticker `json:"price"`
}

// HandleGetStockRequest　株式データ取得
// (GET /stock)
func HandleGetStockRequest(w http.ResponseWriter, r *http.Request, params AP01GetStockParams, u *usecase.StockUsecase) error {

	if params.Symbol == nil || *params.Symbol == "" {
		return fmt.Errorf("")
	}

	price, err := u.GetPrice(r.Context(), *params.Symbol)
	if err != nil {
		return err
	}

	return render.JSONResponse(w, http.StatusAccepted, getStockResponce{Symbol: *params.Symbol, Price: price})
}

package handler

import (
	"encoding/json"
	"go_stock_analysis/domain/model"
	"go_stock_analysis/render"
	"go_stock_analysis/usecase"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type postStockResponce struct {
	Symbol string         `json:"symbol"`
	Price  []model.Ticker `json:"price"`
}

// HandlePostStockRequest　株式データ登録
// (POST /stock)
func HandlePostStockRequest(w http.ResponseWriter, r *http.Request, u *usecase.StockUsecase) error {

	var request AP01PostStockJSONBody
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return err
	}

	if err := request.validate(); err != nil {
		return err
	}

	price, err := u.ResistPriceFromExternalAPI(r.Context(), request.Symbol, request.Interval)
	if err != nil {
		return err
	}

	return render.JSONResponse(w, http.StatusAccepted, postStockResponce{Symbol: request.Symbol, Price: price})
}

func (p AP01PostStockJSONBody) validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Symbol, validation.Required),
		validation.Field(&p.Interval, validation.Required),
	)
}

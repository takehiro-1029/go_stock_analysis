package api

import (
	"go_stock_analysis/infra"
	"go_stock_analysis/registry"
	"go_stock_analysis/server/api/handler"
	"go_stock_analysis/usecase"
	"net/http"

	"github.com/go-chi/chi"
)

// serverImpl Implementation of ServerInterface generated by oapi-codegen command.
type serverImpl struct {
	db registry.Registry
}

// 株価取得
// (GET /stock)
func (s *serverImpl) AP01GetStock(w http.ResponseWriter, r *http.Request, params handler.AP01GetStockParams) {
	sr := infra.NewStockRepository(s.db.DB())
	tr := infra.NewTickerRepository(s.db.DB())
	u := usecase.NewStockUsecase(sr, nil, tr)
	if err := handler.HandleGetStockRequest(w, r, params, u); err != nil {
		msg := err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
	}
}

// 株式データ登録
// (POST /stock)
func (s *serverImpl) AP01PostStock(w http.ResponseWriter, r *http.Request) {
	sr := infra.NewStockRepository(s.db.DB())
	ir := infra.NewIntervalRepository(s.db.DB())
	tr := infra.NewTickerRepository(s.db.DB())
	u := usecase.NewStockUsecase(sr, ir, tr)
	if err := handler.HandlePostStockRequest(w, r, u); err != nil {
		msg := err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
	}
}

// RegisterHandlers APIのハンドラを登録する
func RegisterHandlers(registry registry.Registry, router *chi.Mux) {
	handler.HandlerFromMux(createServerInterface(registry), router)
}

func createServerInterface(registry registry.Registry) handler.ServerInterface {
	return &serverImpl{
		db: registry,
	}
}

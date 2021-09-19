package server

import (
	"compress/flate"
	"go_stock_analysis/registry"
	"go_stock_analysis/server/api"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

const (
	latestAppVersionHeader = "X-Latest-Version"
)

var flateCompressTargets = []string{
	// text: https://www.iana.org/assignments/media-types/media-types.xhtml#text
	"text/html",
	"text/css",
	"text/plain",
	"text/csv",
	"text/xml",
	// application: https://www.iana.org/assignments/media-types/media-types.xhtml#application
	"application/javascript",
	"application/json",
}

// NewServer 新しい一サーバーを作成する
func NewServer(registry registry.Registry, allowCors bool) http.Handler {
	// Setup Router middleware
	r := chi.NewRouter()
	setupMiddleware(r, allowCors)
	api.RegisterHandlers(registry, r)
	return r
}

func setupMiddleware(r chi.Router, allowCors bool) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(flate.DefaultCompression, flateCompressTargets...))
	if allowCors {
		r.Use(cors.Handler(cors.Options{
			AllowOriginFunc:  allowAllOrigin,
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-App-Version", "X-Nonce"},
			ExposedHeaders:   []string{"X-Require-Version", "X-Latest-Version"},
			AllowCredentials: false,
			MaxAge:           300,
		}))
	}
}

func allowAllOrigin(_ *http.Request, origin string) bool {
	return true
}

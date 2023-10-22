package http

import (
	"encoding/json"
	"github.com/brendsanchez/ws-money-go/internal/app"
	"github.com/brendsanchez/ws-money-go/internal/dto"
	"github.com/sirupsen/logrus"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func MapDollarRoutes(handler *http.ServeMux, handlers app.Handlers) {
	handler.HandleFunc("/v1/dollar", handlers.GetPrices())
	handler.HandleFunc("/v1/pages", handlers.GetPages())
}

func MapManageRoutes(handler *http.ServeMux) {
	logrus.Info("health path: ", "/manage/health")
	handler.HandleFunc("/manage/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		r.Header.Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dto.Health{
			Status: "OK",
		})
	})
}

func SwaggerRoute(handler *http.ServeMux) {
	logrus.Info("swagger path: ", "/swagger-ui/")
	handler.Handle("/swagger-ui/", httpSwagger.WrapHandler)
}

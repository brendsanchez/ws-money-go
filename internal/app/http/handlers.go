package http

import (
	"encoding/json"
	"github.com/brendsanchez/ws-money-go/internal/app"
	"github.com/sirupsen/logrus"
	"net/http"
)

type moneyHandlers struct {
	uc app.UseCase
}

// NewMoneyHandlers devmat handlers constructor
func NewMoneyHandlers(uc app.UseCase) app.Handlers {
	return &moneyHandlers{uc: uc}
}

// GetPrices ... Get dollar prices
// @Summary Consigue los tipos de cambio del dolar
// @Description get dollars by param web
// @Tags dollar
// @Accept  json
// @Produce  json
// @Query web path string true page web example(DOLAR_HOY)
// @Success 200  {object}  []dto.Price
// @Failure 400,404,500
// @Router /v1/dollar [get]
func (mh *moneyHandlers) GetPrices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "application/json")

		logrus.Info("getting prices by: ", 12345)

		response := mh.uc.GetPrices("DOLAR_HOY")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
	}
}

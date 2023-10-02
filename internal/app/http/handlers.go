package http

import (
	"encoding/json"
	"github.com/brendsanchez/ws-money-go/internal/app"
	"github.com/brendsanchez/ws-money-go/internal/dto"
	"github.com/sirupsen/logrus"
	"go/types"
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

		webPageParam := r.URL.Query().Get("web")
		if webPageParam == "" {
			writeResponse(w, "query param 'web' is required", http.StatusNotFound)
			return
		}

		logrus.Info("getting prices by: ", webPageParam)
		response := mh.uc.GetPrices(webPageParam)

		w.Header().Set("Content-Type", "application/json")

		responseJSON, err := json.Marshal(response)
		if err != nil {
			writeResponse(w, "Failed to marshal JSON", http.StatusInternalServerError)
			return
		}

		_, err = w.Write(responseJSON)
		if err != nil {
			writeResponse(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	}
}

func writeResponse(w http.ResponseWriter, message string, statusCode int) {
	logrus.Debug(message, statusCode)

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(dto.DollarResponse[types.Nil]{
		Code:    statusCode,
		Message: message,
	})
}

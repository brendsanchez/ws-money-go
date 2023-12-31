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
// @Param web query string true "pages to ws" example(DOLAR_HOY)
// @Success 200 {object} dto.DollarResponse[dto.Dollar]
// @Failure 400,404,500
// @Router /v1/dollar [get]
func (mh *moneyHandlers) GetPrices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "application/json")

		webPageParam := r.URL.Query().Get("web")
		if webPageParam == "" {
			writeResponseError(w, "query param 'web' is required ", http.StatusNotFound)
			return
		}

		logrus.Info("getting prices by: ", webPageParam)
		response := mh.uc.GetPrices(webPageParam)

		w.Header().Set("Content-Type", "application/json")

		responseJSON, err := json.Marshal(response)
		if err != nil {
			writeResponseError(w, "Failed to marshal JSON", http.StatusInternalServerError)
			return
		}

		_, err = w.Write(responseJSON)
		if err != nil {
			writeResponseError(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	}
}

// GetPages ... Get dollar pages
// @Summary Consigue las paginas donde traera los precios
// @Description get pages
// @Tags dollar
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.DollarResponse[dto.Page]
// @Failure 400,404,500
// @Router /v1/pages [get]
func (mh *moneyHandlers) GetPages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "application/json")

		response := mh.uc.GetPages()

		w.Header().Set("Content-Type", "application/json")

		responseJSON, err := json.Marshal(response)
		if err != nil {
			writeResponseError(w, "Failed to marshal JSON", http.StatusInternalServerError)
			return
		}

		_, err = w.Write(responseJSON)
		if err != nil {
			writeResponseError(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	}
}

func writeResponseError(w http.ResponseWriter, message string, statusCode int) {
	logrus.Error(message, statusCode)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(dto.DollarResponse[types.Nil]{
		Code:    statusCode,
		Message: message,
	})
}

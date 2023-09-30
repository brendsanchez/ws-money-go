package scraping

import (
	"errors"
	"github.com/brendsanchez/ws-money-go/internal/dto"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func visitRoute(route string, c *colly.Collector) error {
	c.OnError(func(r *colly.Response, err error) {
		logrus.Error("Request URL:", r.Request.URL, "failed with response:", r, "Error:", err)
	})

	return c.Visit(route)
}

func getTimestamp(text string) *time.Time {
	date := strings.Replace(text, "Actualizado el ", "", 1)
	resul, err := time.Parse("02/01/06 03:04 PM", date)
	if err != nil {
		logrus.Errorf("error parse date dolar hoy: %s", date)
		return nil
	}
	return &resul
}

func getDollarTypes(dollarTypes []dto.Dollar) (*[]dto.Dollar, error) {
	if len(dollarTypes) < 1 {
		return nil, errors.New("not found dollars values")
	}
	return &dollarTypes, nil
}

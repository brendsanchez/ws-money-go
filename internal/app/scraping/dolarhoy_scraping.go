package scraping

import (
	"fmt"
	"github.com/brendsanchez/ws-money-go/internal/app"
	"github.com/brendsanchez/ws-money-go/internal/app/util"
	"github.com/brendsanchez/ws-money-go/internal/dto"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type dolarHoyWS struct {
	route string
}

func NewDolarHoyWS(route string) app.Dollar {
	return &dolarHoyWS{route: route}
}

func (hc *dolarHoyWS) GetPrices() (*[]dto.Dollar, error) {
	c := colly.NewCollector()

	var updatedTime time.Time
	c.OnHTML("div.tile.update", func(el *colly.HTMLElement) {
		updatedTime = getTimestamp(el.Text)
	})

	dollarTypes := make([]dto.Dollar, 0, 6)
	c.OnHTML("div.tile.is-parent.is-7.is-vertical", func(e *colly.HTMLElement) {
		e.ForEach("div.tile.is-child", func(i int, el *colly.HTMLElement) {
			priceBuy := el.ChildText("div.compra div.val")
			priceSell := el.ChildText("div.venta div.val")

			dollar := dto.Dollar{
				Name:      el.ChildText("a"),
				Buy:       &dto.Price{Val: util.ConvertToFloat(priceBuy), ValText: priceBuy},
				Sell:      &dto.Price{Val: util.ConvertToFloat(priceSell), ValText: priceSell},
				Timestamp: updatedTime,
			}
			dollarTypes = append(dollarTypes, dollar)
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "Error:", err)
	})

	err := c.Visit(hc.route)
	if err != nil {
		return nil, err
	}

	return &dollarTypes, nil
}

func getTimestamp(text string) time.Time {
	date := strings.Replace(text, "Actualizado el ", "", 1)

	resul, err := time.Parse("02/01/06 03:04 PM", date)
	if err != nil {
		logrus.Error("error parse date")
		return time.Now()
	}

	return resul
}

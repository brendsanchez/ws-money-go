package scraping

import (
	"errors"
	"github.com/brendsanchez/ws-money-go/internal/app"
	"github.com/brendsanchez/ws-money-go/internal/app/util"
	"github.com/brendsanchez/ws-money-go/internal/dto"
	"github.com/gocolly/colly"
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

	var updatedTime *time.Time
	c.OnHTML("div.tile.update", func(el *colly.HTMLElement) {
		updatedTime = getTimestamp(el.Text)
	})

	dollarTypes := make([]dto.Dollar, 0, 6)
	c.OnHTML("div.tile.is-parent.is-7.is-vertical", func(e *colly.HTMLElement) {
		cant := 1
		e.ForEach("div.tile.is-child", func(i int, el *colly.HTMLElement) {
			priceBuy := el.ChildText("div.compra div.val")
			priceSell := el.ChildText("div.venta div.val")
			dollar := dto.Dollar{
				Id:        cant,
				Name:      el.ChildText("a"),
				Buy:       &dto.Price{Val: util.ConvertToFloat(priceBuy), ValText: priceBuy},
				Sell:      &dto.Price{Val: util.ConvertToFloat(priceSell), ValText: priceSell},
				Timestamp: updatedTime,
			}
			dollarTypes = append(dollarTypes, dollar)
			cant++
		})
	})

	err := visitRoute(hc.route, c)
	if err != nil {
		return nil, errors.New("error visit dolar_hoy")
	}
	return getDollarTypes(dollarTypes)
}

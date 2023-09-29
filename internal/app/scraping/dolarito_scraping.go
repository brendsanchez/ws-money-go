package scraping

import (
	"github.com/brendsanchez/ws-money-go/internal/app"
	"github.com/brendsanchez/ws-money-go/internal/dto"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
)

type dolaritoWS struct {
	route string
}

func NewDolaritoWS(route string) app.Dollar {
	return &dolaritoWS{route: route}
}

func (hc *dolaritoWS) GetPrices() (*[]dto.Dollar, error) {
	c := colly.NewCollector()

	dollarTypes := make([]dto.Dollar, 0, 6)
	c.OnHTML("div.tile.is-parent.is-7.is-vertical", func(e *colly.HTMLElement) {
		e.ForEach("div.tile.is-child", func(i int, el *colly.HTMLElement) {
			dollarTypes = append(dollarTypes, dto.Dollar{
				Name: el.ChildText("a"),
				Buy: dto.Price{Val: el.ChildText("div.compra div.val"),
					Label: el.ChildText("div.compra div.label")},
				Sell: dto.Price{Val: el.ChildText("div.venta div.val"),
					Label: el.ChildText("div.venta div.label")},
			})
		})
		logrus.Debug("RESULT:", dollarTypes)
	})

	c.OnError(func(r *colly.Response, err error) {
		logrus.Error("Request URL: ", r.Request.URL, "failed with response:", r, "Error:", err)
	})

	err := c.Visit(hc.route)
	if err != nil {
		return nil, err
	}

	return &dollarTypes, nil
}

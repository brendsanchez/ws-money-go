package scraping

import (
	"encoding/json"
	"errors"
	"github.com/brendsanchez/ws-money-go/internal/app"
	"github.com/brendsanchez/ws-money-go/internal/app/util"
	"github.com/brendsanchez/ws-money-go/internal/dto"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
	"time"
)

type Dolarito struct {
	Props Props `json:"props"`
}

type Props struct {
	PageProps PageProps `json:"pageProps"`
}

type PageProps struct {
	RealTimeQuotations RealTimeQuotations `json:"realTimeQuotations"`
}

type RealTimeQuotations struct {
	Quotations map[string]Quotation `json:"quotations"`
}

type Quotation struct {
	Name      string `json:"name"`
	Buy       string `json:"buy"`
	Sell      string `json:"sell"`
	Timestamp int64  `json:"timestamp"`
	Variation string `json:"variation"`
	Spread    string `json:"spread"`
	Volume    string `json:"volumen"`
}

type dolaritoWS struct {
	route string
}

func NewDolaritoWS(route string) app.Dollar {
	return &dolaritoWS{route: route}
}

func (hc *dolaritoWS) GetPrices() (*[]dto.Dollar, error) {
	c := colly.NewCollector()

	dollarTypes := make([]dto.Dollar, 0, 16)

	c.OnHTML("script#__NEXT_DATA__", func(e *colly.HTMLElement) {
		jsonData := e.Text
		var data Dolarito

		if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
			logrus.Error("Error al decodificar JSON:", err)
		}

		// Ahora puedes acceder a la lista de "quotations" y sus elementos
		for _, quotation := range data.Props.PageProps.RealTimeQuotations.Quotations {
			timestamp := time.Unix(quotation.Timestamp, 0)
			dollar := dto.Dollar{
				Name:      quotation.Name,
				Sell:      &dto.Price{Val: util.ConvertToFloat(quotation.Sell), ValText: util.AddDollarChar(quotation.Sell)},
				Buy:       &dto.Price{Val: util.ConvertToFloat(quotation.Buy), ValText: util.AddDollarChar(quotation.Buy)},
				Timestamp: &timestamp,
			}
			dollarTypes = append(dollarTypes, dollar)
		}
	})

	err := visitRoute(hc.route, c)
	if err != nil {
		return nil, errors.New("error visit dolarito")
	}
	return getDollarTypes(dollarTypes)
}

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
	Name      string  `json:"name"`
	Buy       float64 `json:"buy"`
	Sell      float64 `json:"sell"`
	Timestamp int64   `json:"timestamp"`
	Variation float64 `json:"variation"`
	Spread    float64 `json:"spread"`
	Volume    float64 `json:"volumen"`
}

type dolaritoWS struct {
	route string
}

func NewDolaritoWS(route string) app.Dollar {
	return &dolaritoWS{route: route}
}

func (do *dolaritoWS) GetPrices() (*[]dto.Dollar, error) {
	c := colly.NewCollector()

	dollarTypes := make([]dto.Dollar, 0)

	c.OnHTML("script#__NEXT_DATA__", func(e *colly.HTMLElement) {
		jsonData := e.Text
		var data Dolarito

		if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
			logrus.Error("Error al decodificar JSON:", err)
			return
		}

		var dateQuotation *time.Time
		for _, quotation := range data.Props.PageProps.RealTimeQuotations.Quotations {
			if quotation.Name != "" {
				valTextSell := util.ConvertValText(quotation.Sell)
				valTextBuy := util.ConvertValText(quotation.Buy)
				dateQuotation = unixToTimestamp(quotation.Timestamp)
				dollar := dto.Dollar{
					Name:      util.CapitalizeWords(quotation.Name),
					Sell:      &dto.Price{Value: util.ConvertToFloat(valTextSell), ValueText: valTextSell},
					Buy:       &dto.Price{Value: util.ConvertToFloat(valTextBuy), ValueText: valTextBuy},
					Date:      dateQuotation,
					Timestamp: quotation.Timestamp,
				}
				dollarTypes = append(dollarTypes, dollar)
			}
		}
	})

	err := visitRoute(do.route, c)
	if err != nil {
		return nil, errors.New("error visit dolarito")
	}
	return &dollarTypes, nil
}

func unixToTimestamp(value int64) *time.Time {
	timestampSeconds := value / 1000
	timestampNanoseconds := (value % 1000) * 1e6
	unix := time.Unix(timestampSeconds, timestampNanoseconds)
	return &unix
}

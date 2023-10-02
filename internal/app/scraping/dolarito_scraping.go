package scraping

import (
	"encoding/json"
	"errors"
	"github.com/brendsanchez/ws-money-go/internal/app"
	"github.com/brendsanchez/ws-money-go/internal/app/util"
	"github.com/brendsanchez/ws-money-go/internal/dto"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
	"strings"
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

func (do *dolaritoWS) GetPrices() (*[]dto.Dollar, error) {
	c := colly.NewCollector()

	dollarTypes := make([]dto.Dollar, 0, 16)

	c.OnHTML("script#__NEXT_DATA__", func(e *colly.HTMLElement) {
		jsonData := e.Text
		var data Dolarito

		if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
			logrus.Error("Error al decodificar JSON:", err)
			return
		}

		var timestamp *time.Time
		for _, quotation := range data.Props.PageProps.RealTimeQuotations.Quotations {
			valTextSell := convertValText(quotation.Sell)
			valTextBuy := convertValText(quotation.Buy)
			timestamp = unixToTimestamp(quotation.Timestamp)
			dollar := dto.Dollar{
				Name:      quotation.Name,
				Sell:      &dto.Price{Val: util.ConvertToFloat(valTextSell), ValText: valTextSell},
				Buy:       &dto.Price{Val: util.ConvertToFloat(valTextBuy), ValText: valTextBuy},
				Timestamp: timestamp,
			}
			dollarTypes = append(dollarTypes, dollar)
		}
	})

	err := visitRoute(do.route, c)
	if err != nil {
		return nil, errors.New("error visit dolarito")
	}
	return &dollarTypes, nil
}

func convertValText(value string) string {
	if value == "" {
		return value
	}
	return "$" + strings.Replace(value, ",", ".", 1)
}

func unixToTimestamp(value int64) *time.Time {
	timestampSeconds := value / 1000
	timestampNanoseconds := (value % 1000) * 1e6
	unix := time.Unix(timestampSeconds, timestampNanoseconds)
	return &unix
}

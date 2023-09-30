package scraping

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/brendsanchez/ws-money-go/internal/app"
	"github.com/brendsanchez/ws-money-go/internal/dto"
	"github.com/gocolly/colly"
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

// Define el struct para representar un elemento de "quotations"
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
	var dolarito Dolarito

	// Visita la URL y analiza la respuesta
	c.OnHTML("script#__NEXT_DATA__", func(e *colly.HTMLElement) {
		// Obtiene el contenido del elemento <script> con el ID "__NEXT_DATA__"
		scriptContent := e.Text

		// Decodifica el contenido JSON directamente en el struct quotationsList
		if err := json.Unmarshal([]byte(scriptContent), &dolarito); err != nil {
			fmt.Println("Error al decodificar JSON:", err)
		}

		// Ahora puedes acceder a la lista de "quotations" y sus elementos
		for _, quotation := range dolarito.Props.PageProps.RealTimeQuotations.Quotations {
			fmt.Println("Nombre:", quotation.Name)
			fmt.Println("Cotización de compra:", quotation.Buy)
			fmt.Println("Cotización de venta:", quotation.Sell)
			// Accede a otros campos según sea necesario
		}
	})

	err := visitRoute(hc.route, c)
	if err != nil {
		return nil, errors.New("error visit dolarito")
	}
	return getDollarTypes(dollarTypes)
}

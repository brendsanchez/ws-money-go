package scraping

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/brendsanchez/ws-money-go/internal/app"
	"github.com/brendsanchez/ws-money-go/internal/dto"
	"github.com/gocolly/colly"
)

// Define el struct para representar la lista de "quotations"
type QuotationsList struct {
	Quotations []Quotation `json:"quotations"`
}

// Define el struct para representar un elemento de "quotations"
type Quotation struct {
	Name      string            `json:"name"`
	Buy       string            `json:"buy"`
	Sell      string            `json:"sell"`
	Timestamp int64             `json:"timestamp"`
	Variation string            `json:"variation"`
	Spread    string            `json:"spread"`
	Volume    map[string]string `json:"volumen"`
	Extra     ExtraInfo         `json:"extra"`
}

// Define el struct para representar la información adicional en "extra"
type ExtraInfo struct {
	ReferenceBuy1 struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	} `json:"referenceBuy1"`

	ReferenceBuy2 struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	} `json:"referenceBuy2"`
}

type dolaritoWS struct {
	route string
}

func NewDolaritoWS(route string) app.Dollar {
	return &dolaritoWS{route: route}
}

func (hc *dolaritoWS) GetPrices() (*[]dto.Dollar, error) {
	c := colly.NewCollector()

	dollarTypes := make([]dto.Dollar, 0, 6)
	var quotationsList QuotationsList

	// Visita la URL y analiza la respuesta
	c.OnHTML("script#__NEXT_DATA__", func(e *colly.HTMLElement) {
		// Obtiene el contenido del elemento <script> con el ID "__NEXT_DATA__"
		scriptContent := e.Text

		// Decodifica el contenido JSON directamente en el struct quotationsList
		if err := json.Unmarshal([]byte(scriptContent), &quotationsList); err != nil {
			fmt.Println("Error al decodificar JSON:", err)
		}

		// Ahora puedes acceder a la lista de "quotations" y sus elementos
		for _, quotation := range quotationsList.Quotations {
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

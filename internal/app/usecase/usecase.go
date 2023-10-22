package usecase

import (
	"github.com/brendsanchez/ws-money-go/config"
	"github.com/brendsanchez/ws-money-go/internal/app/factory"
	"github.com/brendsanchez/ws-money-go/internal/dto"
	"github.com/brendsanchez/ws-money-go/internal/enums"
	"github.com/sirupsen/logrus"
	"net/http"
	"sort"
)

type dollarUC struct {
	cfg *config.Config
}

func NewDollarUC(cfg *config.Config) *dollarUC {
	return &dollarUC{cfg: cfg}
}

func (duc *dollarUC) GetPrices(webPage string) dto.DollarResponse[*[]dto.Dollar] {
	enumPageWeb := enums.WebPage(webPage)

	moneyFactory := factory.NewMoneyFactory(duc.cfg)
	dollar, err := moneyFactory.GetFactory(enumPageWeb)
	if err != nil {
		logrus.Error(err.Error())
		return dollarResponse[*[]dto.Dollar](err.Error(), http.StatusNotFound, nil)
	}

	pricesList, err := dollar.GetPrices()
	if err != nil {
		logrus.Errorf("error scraping %s", err.Error())
		return dollarResponse[*[]dto.Dollar](err.Error(), http.StatusInternalServerError, nil)
	}

	if len(*pricesList) < 1 {
		logrus.Error("not found dollars values", enumPageWeb)
		return dollarResponse[*[]dto.Dollar]("not found dollars values", http.StatusNotFound, nil)
	}

	sort.Sort(dto.ByName(*pricesList))
	return dollarResponse("success", http.StatusOK, pricesList)
}

func (duc *dollarUC) GetPages() dto.DollarResponse[*[]dto.Page] {
	routes := duc.cfg.Routes
	if len(routes) < 1 {
		logrus.Error("error al conseguir las paginas")
		return dollarResponse[*[]dto.Page]("pages not found", http.StatusNotFound, nil)
	}

	var pages []dto.Page
	for _, route := range routes {
		pages = append(pages, dto.Page{Name: route.Name, Route: route.Uri})
	}

	return dollarResponse("success", http.StatusOK, &pages)
}

func dollarResponse[T any](message string, code int, data T) dto.DollarResponse[T] {
	return dto.DollarResponse[T]{
		Message: message,
		Code:    code,
		Data:    data,
	}
}

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
		return dollarResponse(err.Error(), http.StatusNotFound, nil)
	}

	pricesList, err := dollar.GetPrices()
	if err != nil {
		logrus.Errorf("error scraping %s", err.Error())
		return dollarResponse(err.Error(), http.StatusInternalServerError, nil)
	}

	if len(*pricesList) < 1 {
		logrus.Error("not found dollars values", enumPageWeb)
		return dollarResponse("not found dollars values", http.StatusNotFound, nil)
	}

	sort.Sort(dto.ByName(*pricesList))
	return dollarResponse("success", http.StatusOK, pricesList)
}

func dollarResponse(message string, code int, data *[]dto.Dollar) dto.DollarResponse[*[]dto.Dollar] {
	return dto.DollarResponse[*[]dto.Dollar]{
		Message: message,
		Code:    code,
		Data:    data}
}

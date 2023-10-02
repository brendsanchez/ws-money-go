package usecase

import (
	"github.com/brendsanchez/ws-money-go/config"
	"github.com/brendsanchez/ws-money-go/internal/app/factory"
	"github.com/brendsanchez/ws-money-go/internal/dto"
	"github.com/brendsanchez/ws-money-go/internal/enums"
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
		return dollarResponse(err.Error(), http.StatusNotFound, nil)
	}

	pricesList, err := dollar.GetPrices()
	if err != nil {
		return dollarResponse(err.Error(), http.StatusNotFound, nil)
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

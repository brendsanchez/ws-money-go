package usecase

import (
	"github.com/brendsanchez/ws-money-go/config"
	"github.com/brendsanchez/ws-money-go/internal/app/factory"
	"github.com/brendsanchez/ws-money-go/internal/dto"
	"github.com/brendsanchez/ws-money-go/internal/enums"
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
		return dollarResponseNotFound(err)
	}

	pricesList, err := dollar.GetPrices()
	if err != nil {
		return dollarResponseNotFound(err)
	}

	sort.Sort(dto.ByName(*pricesList))
	
	return dto.DollarResponse[*[]dto.Dollar]{
		Message: "success",
		Code:    200,
		Data:    pricesList}
}

func dollarResponseNotFound(err error) dto.DollarResponse[*[]dto.Dollar] {
	return dto.DollarResponse[*[]dto.Dollar]{
		Message: err.Error(),
		Code:    404,
		Data:    nil}
}

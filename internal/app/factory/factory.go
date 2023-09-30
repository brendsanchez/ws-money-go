package factory

import (
	"errors"
	"github.com/brendsanchez/ws-money-go/config"
	"github.com/brendsanchez/ws-money-go/internal/app"
	"github.com/brendsanchez/ws-money-go/internal/app/scraping"
	"github.com/brendsanchez/ws-money-go/internal/enums"
	"github.com/sirupsen/logrus"
)

type dollarFactory struct {
	cfg *config.Config
}

func NewMoneyFactory(cfg *config.Config) *dollarFactory {
	return &dollarFactory{cfg: cfg}
}

func (f *dollarFactory) GetFactory(pageWeb enums.WebPage) (app.Dollar, error) {
	logrus.Debugf("pageWeb: %s", pageWeb)

	if pageWeb == enums.DOLAR_HOY {
		return scraping.NewDolarHoyWS(f.cfg.Route.DolarHoy), nil
	}

	if pageWeb == enums.DOLARITO {
		return scraping.NewDolaritoWS(f.cfg.Route.Dolarito), nil
	}
	return nil, errors.New("not found page web valid")
}

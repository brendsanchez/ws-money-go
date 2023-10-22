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
	if pageWeb == enums.DOLAR_HOY {
		logrus.Debug("choosing factory dolar hoy")
		return scraping.NewDolarHoyWS(f.cfg.Routes[0].Uri), nil
	}
	if pageWeb == enums.DOLARITO {
		logrus.Debug("choosing factory dolarito")
		return scraping.NewDolaritoWS(f.cfg.Routes[1].Uri), nil
	}
	return nil, errors.New("not found web=" + string(pageWeb))
}

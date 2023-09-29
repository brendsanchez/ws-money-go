package server

import (
	"github.com/brendsanchez/ws-money-go/docs"
	"github.com/brendsanchez/ws-money-go/internal/app/http"
	"github.com/brendsanchez/ws-money-go/internal/app/usecase"
	"github.com/sirupsen/logrus"
)

func (s *Server) mapHandlers() {

	// Init useCase
	logrus.Debug("init useCase")
	dolarUseCase := usecase.NewDollarUC(s.cfg)

	// Init handlers
	logrus.Debug("init handlers")
	dollarHandlers := http.NewMoneyHandlers(dolarUseCase)

	// Swagger
	docs.SwaggerInfo.Version = s.cfg.App.Version

	// Routes
	logrus.Debug("init routes")
	http.SwaggerRoute(s.handler)
	http.MapManageRoutes(s.handler)
	http.MapDollarRoutes(s.handler, dollarHandlers)
}

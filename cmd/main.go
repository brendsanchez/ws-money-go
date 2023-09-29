package main

import (
	"github.com/brendsanchez/ws-money-go/config"
	"github.com/brendsanchez/ws-money-go/internal/server"
	"github.com/sirupsen/logrus"
)

// @title ws-money
// @description cotizaciones del dolar
// @contact.name dusk
// @contact.email brendasanchez9310@gmail.com
// @externalDocs.description github
// @externalDocs.url https://github.com/brendsanchez
func main() {
	cfg, err := config.Load()
	if err != nil {
		logrus.Fatal("Error config")
	}

	config.InitLogrus(cfg)
	sv := server.NewServer(cfg)
	if sv.Run() != nil {
		logrus.Fatal("Error server")
	}
}

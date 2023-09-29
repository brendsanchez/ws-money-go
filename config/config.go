package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
	"time"
)

type Config struct {
	App    app
	Route  route
	Server server
	Logger logger
}

type app struct {
	Name    string
	Version string
	Mode    string `env:"APP_MODE, default=dev"`
}

type route struct {
	Dolarito string
	DolarHoy string
}

type server struct {
	Port           string        `env:"SERVER_PORT, default=8080"`
	ReadTimeout    time.Duration `env:"SERVER_READ_TIMEOUT, default=20s"`
	WriteTimeout   time.Duration `env:"SERVER_WRITE_TIMEOUT, default=20s"`
	IdleTimeout    time.Duration `env:"SERVER_IDLE_TIMEOUT, default=120s"`
	MaxHeaderBytes int           `env:"SERVER_MAX_HEADER_BYTES, default=32768"`
}

type logger struct {
	IsDebug      bool `env:"LOG_IS_DEBUG, default=true"`
	ReportCaller bool `env:"LOG_REPORT_CALLER, default=false"`
}

func Load() (*Config, error) {
	ctx := context.Background()
	var config Config
	if err := envconfig.Process(ctx, &config); err != nil {
		return nil, err
	}
	config.App.Name = "ws-money"
	config.App.Version = "1.0.0"
	config.Route.Dolarito = "https://www.dolarito.ar/"
	config.Route.DolarHoy = "https://dolarhoy.com/"
	return &config, nil
}

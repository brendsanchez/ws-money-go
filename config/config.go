package config

import (
	"context"
	"github.com/brendsanchez/ws-money-go/internal/enums"
	"github.com/sethvargo/go-envconfig"
	"os"
	"time"
)

type Config struct {
	App    app
	Routes []route
	Server server
	Logger logger
}

type app struct {
	Name    string
	Version string
	Mode    string `env:"APP_MODE, default=dev"`
}

type route struct {
	Id       string
	Name     string
	Uri      string
	ImageUri string
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
	config.Routes = []route{
		{Id: string(enums.DOLAR_HOY), Name: "Dolar Hoy", Uri: "https://dolarhoy.com/", ImageUri: os.Getenv("DOLAR_HOY_IMG")},
		{Id: string(enums.DOLARITO), Name: "Dolarito", Uri: "https://www.dolarito.ar/", ImageUri: os.Getenv("DOLARITO_IMG")},
	}
	return &config, nil
}

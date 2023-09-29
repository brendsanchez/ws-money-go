package server

import (
	"context"
	"github.com/brendsanchez/ws-money-go/config"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	ctxTimeout = 5
)

// Server struct
type Server struct {
	handler *http.ServeMux
	cfg     *config.Config
}

// NewServer New Server constructor
func NewServer(cfg *config.Config) *Server {
	return &Server{handler: http.NewServeMux(), cfg: cfg}
}

func (s *Server) Run() error {
	logger := logrus.New()
	w := logger.Writer()
	defer w.Close()

	s.mapHandlers()
	server := &http.Server{
		Addr:           ":" + s.cfg.Server.Port,
		Handler:        s.handler,
		ReadTimeout:    s.cfg.Server.ReadTimeout,
		WriteTimeout:   s.cfg.Server.WriteTimeout,
		IdleTimeout:    s.cfg.Server.IdleTimeout,
		MaxHeaderBytes: s.cfg.Server.MaxHeaderBytes,
		ErrorLog:       log.New(w, "server", 0),
	}

	go func() {
		logrus.Infof("Starting App, name: %s ,version: %s, mode: %s", s.cfg.App.Name, s.cfg.App.Version, s.cfg.App.Mode)
		logrus.Debug("Config:", s.cfg)
		logrus.Infof("Server is listening port: %s", s.cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil {
			logrus.Fatal("Error server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	return server.Shutdown(ctx)
}

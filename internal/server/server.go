// Package server provides HTTP server management.
package server

import (
	"context"
	"net/http"

	"github.com/lipandr/go-chat-micoservice/internal/config"
	"github.com/sirupsen/logrus"
)

// Server wraps HTTP server with logging.
type Server struct {
	httpServer *http.Server
	logger     *logrus.Logger
}

// New creates a new Server instance.
func New(handler http.Handler, cfg config.ServerConfig, appCfg config.AppConfig, logger *logrus.Logger) *Server {
	srv := &Server{
		logger: logger,
	}

	srv.httpServer = &http.Server{
		Addr:         cfg.Address(),
		Handler:      handler,
		ReadTimeout:  appCfg.ReadTimeout,
		WriteTimeout: appCfg.WriteTimeout,
		IdleTimeout:  appCfg.IdleTimeout,
	}

	return srv
}

// Start runs the HTTP server.
func (s *Server) Start() error {
	s.logger.Info("starting server", "addr", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

// Stop gracefully shuts down the HTTP server.
func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("stopping server")
	return s.httpServer.Shutdown(ctx)
}

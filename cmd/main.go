// Package main provides entry point for chat microservice.
package main

import (
	"context"
	"net/http"

	"github.com/lipandr/go-chat-micoservice/internal/config"
	"github.com/lipandr/go-chat-micoservice/internal/logger"
	"github.com/lipandr/go-chat-micoservice/internal/server"
	"github.com/lipandr/go-chat-micoservice/internal/shutdown"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	log := logger.NewDefault()

	cfg, err := config.Load()
	if err != nil {
		log.Error("failed to load config", "error", err)
		return err
	}

	appCfg := config.NewAppConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	srv := server.New(mux, cfg.Server, appCfg, log)

	shutdownManager := shutdown.New(log)
	shutdownManager.AddHook(func(ctx context.Context) error {
		return srv.Stop(ctx)
	})

	go func() {
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Error("server error", "error", err)
		}
	}()

	return shutdownManager.Wait(func() error { return nil })
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}

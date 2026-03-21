// Package shutdown manages graceful application shutdown.
package shutdown

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

// Hook is a function called during shutdown.
type Hook func(context.Context) error

// New creates a new shutdown Manager.
func New(logger *logrus.Logger, hooks ...Hook) *Manager {
	return &Manager{
		logger: logger,
		hooks:  hooks,
	}
}

// Manager orchestrates graceful shutdown.
type Manager struct {
	logger *logrus.Logger
	hooks  []Hook
}

// AddHook registers a new shutdown hook.
func (m *Manager) AddHook(hook Hook) {
	m.hooks = append(m.hooks, hook)
}

// Wait blocks until shutdown signal is received, then runs all hooks.
func (m *Manager) Wait(interrupt func() error) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-quit:
		m.logger.Info("received signal", "signal", sig.String())
	case err := <-errChannel(interrupt):
		if err != nil {
			return err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	for i, hook := range m.hooks {
		if err := hook(ctx); err != nil {
			m.logger.Error("shutdown hook failed", "index", i, "error", err)
			if errors.Is(err, context.DeadlineExceeded) {
				return errors.New("shutdown timeout exceeded")
			}
			return err
		}
	}

	m.logger.Info("shutdown completed")
	return nil
}

func errChannel(f func() error) chan error {
	errCh := make(chan error, 1)
	go func() {
		errCh <- f()
	}()
	return errCh
}

const shutdownTimeout = 30 * time.Second

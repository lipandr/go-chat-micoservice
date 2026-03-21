// Package logger provides structured logging configuration.
package logger

import (
	"github.com/sirupsen/logrus"
)

// New creates a new logrus.Logger with JSON formatter.
func New(level string) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(lvl)
	}

	return logger
}

// NewDefault creates a logger with info level.
func NewDefault() *logrus.Logger {
	return New("info")
}

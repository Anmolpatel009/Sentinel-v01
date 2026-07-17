// Package logging provides Sentinel's logging subsystem.
//
// Mission:
//
// The Logger is the public face of the Logging subsystem.
//
// Every other package in Sentinel should interact only with this object.
// The rest of Sentinel should never import or use log/slog directly.
package logging

import (
	"io"
	"log/slog"
)

// Logger provides Sentinel's logging API.
//
// It wraps Go's slog.Logger to keep the rest of the codebase
// independent of the underlying logging implementation.
type Logger struct {
	logger *slog.Logger
	config Config
}

// New creates a new Logger instance.
func New(cfg Config, output io.Writer) (*Logger, error) {
	handler, err := NewHandler(cfg, output)
	if err != nil {
		return nil, err
	}

	return &Logger{
		logger: slog.New(handler),
		config: cfg,
	}, nil
}

// Debug logs a debug message.
func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

// Info logs an informational message.
func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

// Warn logs a warning message.
func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

// Error logs an error message.
func (l *Logger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

// With returns a child logger with additional structured attributes.
func (l *Logger) With(args ...any) *Logger {
	return &Logger{
		logger: l.logger.With(args...),
		config: l.config,
	}
}

// Config returns the logger configuration.
func (l *Logger) Config() Config {
	return l.config
}

// Package logging implements the Runtime Component interface.
//
// This file adapts the Logger to Sentinel's runtime.Component
// interface. The Runtime owns lifecycle state, while the Logger
// owns logging behavior.
package logging

import (
	"context"
)

// Name returns the unique component name.
func (l *Logger) Name() string {
	return "logging"
}

// Init prepares the logging subsystem.
func (l *Logger) Init(ctx context.Context) error {
	l.Info("Logging subsystem initialized")
	return nil
}

// Start starts the logging subsystem.
func (l *Logger) Start(ctx context.Context) error {
	l.Info("Logging subsystem started")
	return nil
}

// Stop gracefully shuts down the logging subsystem.
func (l *Logger) Stop(ctx context.Context) error {
	l.Info("Logging subsystem stopping")

	// Future:
	// - Flush buffered logs
	// - Close log files
	// - Shutdown remote log exporters

	return nil
}

// Health reports the health of the logging subsystem.
func (l *Logger) Health(ctx context.Context) error {
	// v1:
	// Logger has no external dependencies,
	// so if it exists, it is considered healthy.
	return nil
}

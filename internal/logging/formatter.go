//formatter.go responsibilities

//├── Convert Config → slog.Handler
//├── Configure JSON/Text handlers
//├── Configure Log Levels
//└── Configure Source Information

//internal flow will be ->
//config->validdate()-> convert log level -> { create slog.handleroptions }  ->  json or text -> return handler

//config-> formatter.go->slog.handler-> logger.go->runtime/componenets

package logging

import (
	"fmt"
	"io"
	"log/slog"
)

// NewHandler creates a slog.Handler based on the provided logging configuration.
// The handler writes to the supplied output writer.
func NewHandler(cfg Config, output io.Writer) (slog.Handler, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	opts := &slog.HandlerOptions{
		Level:     toSlogLevel(cfg.Level),
		AddSource: cfg.AddSource,
	}

	switch cfg.Format {
	case TextFormat:
		return slog.NewTextHandler(output, opts), nil

	case JSONFormat:
		return slog.NewJSONHandler(output, opts), nil

	default:
		return nil, fmt.Errorf("logging: unsupported log format %q", cfg.Format)
	}
}

// toSlogLevel converts Sentinel log levels to slog log levels.
func toSlogLevel(level LogLevel) slog.Level {
	switch level {
	case DebugLevel:
		return slog.LevelDebug

	case InfoLevel:
		return slog.LevelInfo

	case WarnLevel:
		return slog.LevelWarn

	case ErrorLevel:
		return slog.LevelError

	default:
		return slog.LevelInfo
	}
}

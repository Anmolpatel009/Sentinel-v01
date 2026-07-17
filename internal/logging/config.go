//Configuration Types

//↓

//Default Configuration

//↓

//Validation

package logging

import "fmt"

// ============================================================================
// Log Levels
// ============================================================================

// LogLevel represents the severity level of a log message.
type LogLevel string

const (
	// DebugLevel is used for detailed debugging information.
	DebugLevel LogLevel = "DEBUG"

	// InfoLevel is used for normal operational events.
	InfoLevel LogLevel = "INFO"

	// WarnLevel indicates unexpected but recoverable situations.
	WarnLevel LogLevel = "WARN"

	// ErrorLevel indicates failures that require attention.
	ErrorLevel LogLevel = "ERROR"
)

// ============================================================================
// Log Formats
// ============================================================================

// LogFormat defines the output format of the logger.
type LogFormat string

const (
	// TextFormat outputs human-readable logs.
	TextFormat LogFormat = "text"

	// JSONFormat outputs structured JSON logs.
	JSONFormat LogFormat = "json"
)

// ============================================================================
// Logger Configuration
// ============================================================================

// Config defines the configuration for the Sentinel logging subsystem.
type Config struct {

	// Level defines the minimum log level to emit.
	Level LogLevel

	// Format specifies the log output format.
	Format LogFormat

	// AddSource includes source file and line information in logs.
	AddSource bool
}

// DefaultConfig returns the default logging configuration.
func DefaultConfig() Config {
	return Config{
		Level:     InfoLevel,
		Format:    TextFormat,
		AddSource: true,
	}
}

// Validate verifies that the configuration contains valid values.
func (c Config) Validate() error {

	switch c.Level {
	case DebugLevel, InfoLevel, WarnLevel, ErrorLevel:
	default:
		return fmt.Errorf("logging: invalid log level %q", c.Level)
	}

	switch c.Format {
	case TextFormat, JSONFormat:
	default:
		return fmt.Errorf("logging: invalid log format %q", c.Format)
	}

	return nil
}

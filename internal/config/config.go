// Package config provides the configuration model for Sentinel.
//
// Mission:
//
// The configuration subsystem acts as the single source of truth for
// all runtime configuration. No other package should parse configuration
// files directly.
package config

// Config is the root configuration object for Sentinel.
//
// Every subsystem receives its configuration from this object.
// As Sentinel evolves, new subsystem configurations will be added here.
type Config struct {
	Runtime RuntimeConfig `yaml:"runtime"`
	Logging LoggingConfig `yaml:"logging"`
}

// RuntimeConfig contains runtime-specific configuration.
type RuntimeConfig struct {

	// Name identifies this Sentinel instance.
	Name string `yaml:"name"`
}

// LoggingConfig contains logging configuration.
type LoggingConfig struct {

	// Level defines the minimum log level.
	//
	// Supported values:
	//   - DEBUG
	//   - INFO
	//   - WARN
	//   - ERROR
	Level string `yaml:"level"`

	// Format defines the output format.
	//
	// Supported values:
	//   - text
	//   - json
	Format string `yaml:"format"`

	// AddSource includes source file information in log entries.
	AddSource *bool `yaml:"add_source"`
}

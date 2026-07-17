// Package config provides configuration validation for Sentinel.
package config

import (
	"fmt"
	"strings"
)

// Validate checks whether the configuration is valid.
//
// Validation is performed after defaults have been applied.
// All validation errors are collected and returned together.
func Validate(cfg *Config) error {
	var errs []string

	// ----------------------------
	// Runtime Validation
	// ----------------------------

	if strings.TrimSpace(cfg.Runtime.Name) == "" {
		errs = append(errs, "runtime.name must not be empty")
	}

	// ----------------------------
	// Logging Validation
	// ----------------------------

	switch strings.ToUpper(cfg.Logging.Level) {
	case "DEBUG", "INFO", "WARN", "ERROR":
		// Valid
	default:
		errs = append(errs,
			fmt.Sprintf(
				"logging.level '%s' is invalid (allowed: DEBUG, INFO, WARN, ERROR)",
				cfg.Logging.Level,
			),
		)
	}

	switch strings.ToLower(cfg.Logging.Format) {
	case "text", "json":
		// Valid
	default:
		errs = append(errs,
			fmt.Sprintf(
				"logging.format '%s' is invalid (allowed: text, json)",
				cfg.Logging.Format,
			),
		)
	}

	// ----------------------------
	// Final Result
	// ----------------------------

	if len(errs) > 0 {
		return fmt.Errorf("configuration validation failed:\n - %s",
			strings.Join(errs, "\n - "))
	}

	return nil
}

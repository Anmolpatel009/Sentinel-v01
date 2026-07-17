// code flow
/*    sentinel.yaml
                  │
                  ▼
            os.ReadFile()
                  │
                  ▼
           yaml.Unmarshal()
                  │
                  ▼
              Config Struct
                  │
                  ▼
          ApplyDefaults()
                  │
                  ▼
             Validate()
                  │
      ┌───────────┴───────────┐
      ▼                       ▼
  Validation OK        Validation Failed
      │                       │
      ▼                       ▼
Return *Config          Return Error
*/

// Package config provides configuration loading for Sentinel.
//
// In Sentinel v1, configuration is loaded exclusively from a YAML file.
// Future versions may introduce additional configuration sources
// (e.g. ConfigMaps, environment variables, remote APIs) behind a
// decoder abstraction.
package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// DefaultConfigPath is the default location of Sentinel's
// configuration file.
const DefaultConfigPath = "configs/sentinel.yaml"

// Load reads a Sentinel configuration file, unmarshals the YAML,
// applies default values, validates the final configuration,
// and returns a ready-to-use Config instance.
func Load(path string) (*Config, error) {

	// Read the configuration file.
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read configuration file %q: %w", path, err)
	}

	// Parse the YAML into the configuration structure.
	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse configuration file %q: %w", path, err)
	}

	// Fill any unspecified values with Sentinel defaults.
	ApplyDefaults(&cfg)

	// Validate the completed configuration before the runtime starts.
	if err := Validate(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

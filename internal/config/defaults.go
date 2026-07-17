// arch flow of this code ->

/*sentinel.yaml
       │
       ▼
    Load YAML
       │
       ▼
 Partial Config
       │
       ▼
Apply Defaults
       │
       ▼
Complete Config
       │
       ▼
  Validate
       │
       ▼
Runtime


*/

// Package config provides Sentinel's default configuration values.
package config

// ApplyDefaults fills any missing configuration values with
// Sentinel's default settings.
//
// Defaults are applied after loading the configuration file
// and before validation.
func ApplyDefaults(cfg *Config) {

	// ----------------------------
	// Runtime Defaults
	// ----------------------------

	if cfg.Runtime.Name == "" {
		cfg.Runtime.Name = "sentinel-v1"
	}

	// ----------------------------
	// Logging Defaults
	// ----------------------------

	if cfg.Logging.Level == "" {
		cfg.Logging.Level = "INFO"
	}

	if cfg.Logging.Format == "" {
		cfg.Logging.Format = "text"
	}

	// Default: include source information.
	if cfg.Logging.AddSource == nil {
		v := true
		cfg.Logging.AddSource = &v
	}
}

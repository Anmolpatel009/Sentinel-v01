// Package config provides configuration management for Sentinel.
//
// The ConfigManager owns the lifecycle of Sentinel's configuration.
// In Sentinel v1, configuration is loaded once during startup.
// Future versions may support configuration reloads.
package config

// Manager owns the loaded Sentinel configuration.
type Manager struct {
	config *Config
}

// NewManager creates a new configuration manager.
func NewManager() *Manager {
	return &Manager{}
}

// Load loads, defaults, validates, and stores the configuration.
func (m *Manager) Load(path string) error {
	cfg, err := Load(path)
	if err != nil {
		return err
	}

	m.config = cfg
	return nil
}

// Config returns the loaded configuration.
//
// The returned configuration should be treated as read-only.
func (m *Manager) Config() *Config {
	return m.config
}

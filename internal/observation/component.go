package observation

import (
	"context"

	"github.com/Anmolpatel009/Sentinel-v01/internal/runtime"
)

// Component integrates the Observation Engine
// with the Sentinel runtime lifecycle.
type Component struct {
	manager *Manager
	state   runtime.LifecycleState
}

// NewComponent creates a new Observation component.
func NewComponent(manager *Manager) (*Component, error) {
	if manager == nil {
		return nil, ErrManagerNil
	}

	return &Component{
		manager: manager,
		state:   runtime.LifecycleCreated,
	}, nil
}

// Name returns the component name.
func (c *Component) Name() string {
	return "observation"
}

// Init prepares the Observation Engine.
func (c *Component) Init(ctx context.Context) error {
	c.state = runtime.LifecycleInitialized
	return nil
}

// Start starts the Observation Engine.
func (c *Component) Start(ctx context.Context) error {
	c.state = runtime.LifecycleRunning
	return nil
}

// Stop gracefully stops the Observation Engine.
func (c *Component) Stop(ctx context.Context) error {
	c.state = runtime.LifecycleStopping

	c.state = runtime.LifecycleStopped

	return nil
}

// Health reports Observation health.
func (c *Component) Health(ctx context.Context) error {
	if c == nil {
		return ErrComponentNil
	}

	if c.manager == nil {
		return ErrManagerNil
	}

	return nil
}

/*
// Healthy is a convenience wrapper around Health.
func (c *Component) Healthy(ctx context.Context) bool {
	return c.Health(ctx) == nil
}
*/

// State returns the lifecycle state.
func (c *Component) State() runtime.LifecycleState {
	return c.state
}

// Manager exposes the Observation manager.
func (c *Component) Manager() *Manager {
	return c.manager
}

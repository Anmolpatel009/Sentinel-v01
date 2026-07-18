package state

import (
	"context"

	"github.com/Anmolpatel009/Sentinel-v01/internal/runtime"
)

// Component represents the SSR subsystem
// inside Sentinel runtime.
//
// It connects the Shared State Repository
// with the Sentinel lifecycle manager.
type Component struct {
	repository Repository
	state      runtime.LifecycleState
}

// NewComponent creates a new SSR component.
func NewComponent(repository Repository) (*Component, error) {
	if repository == nil {
		return nil, ErrRepositoryNil
	}

	return &Component{
		repository: repository,
		state:      runtime.LifecycleCreated,
	}, nil
}

// Name returns the component identifier.
func (c *Component) Name() string {

	return "state"
}

// Init initializes SSR.
func (c *Component) Init(ctx context.Context) error {

	c.state = runtime.LifecycleInitialized

	return nil
}

// Start starts SSR.
func (c *Component) Start(ctx context.Context) error {

	c.state = runtime.LifecycleRunning

	return nil
}

// Stop gracefully shuts down SSR.
func (c *Component) Stop(ctx context.Context) error {

	c.state = runtime.LifecycleStopping

	// Clear in-memory state.
	if c.repository != nil {

		if err := c.repository.Clear(ctx); err != nil {
			return err
		}
	}

	c.state = runtime.LifecycleStopped

	return nil
}

// Health checks SSR health.
func (c *Component) Health(ctx context.Context) error {
	if c == nil {
		return ErrComponentNil
	}

	if c.repository == nil {
		return ErrRepositoryNil
	}

	return c.repository.Health(ctx)
}

// State returns current lifecycle state.
func (c *Component) State() runtime.LifecycleState {

	return c.state
}

// Repository exposes the SSR repository.
//
// Other components should receive this dependency
// instead of creating their own store.
func (c *Component) Repository() Repository {

	return c.repository
}

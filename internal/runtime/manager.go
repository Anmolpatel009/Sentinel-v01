//What is the Runtime?
//Responsibilities:

//Runtime struct
//Constructor
//Component registry
//Runtime methods

package runtime

import (
	"context"
	"fmt"
	"sync"
)

// Runtime owns the lifecycle of the Sentinel process.
//
// It is responsible for registering, starting, monitoring,
// and stopping all Sentinel components.
//
// Runtime never contains business logic.
// It only orchestrates component lifecycles.
type Runtime struct {
	mu sync.RWMutex

	state LifecycleState

	components map[string]Component
}

// New creates a new Runtime instance.
func New() *Runtime {
	return &Runtime{
		state:      LifecycleCreated,
		components: make(map[string]Component),
	}
}

// Register registers a component with the Runtime.
func (r *Runtime) Register(component Component) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	name := component.Name()

	if _, exists := r.components[name]; exists {
		return fmt.Errorf("runtime: component %q already registered", name)
	}

	r.components[name] = component

	return nil
}

// State returns the current Runtime lifecycle state.
func (r *Runtime) State() LifecycleState {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.state
}

// Health verifies the health of all registered components.
func (r *Runtime) Health(ctx context.Context) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, component := range r.components {
		if err := component.Health(ctx); err != nil {
			return err
		}
	}

	return nil
}

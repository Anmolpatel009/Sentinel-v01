//his file answers:

//How does every component behave?

//It will eventually contain:

//Component interface
//Lifecycle states
//State transitions

package runtime

import "context"

// ============================================================================
// Lifecycle States
// ============================================================================

// LifecycleState represents the current lifecycle state of a Sentinel component.
type LifecycleState int

const (
	// LifecycleUnknown represents an undefined or uninitialized state.
	LifecycleUnknown LifecycleState = iota

	// LifecycleCreated indicates the component has been instantiated.
	LifecycleCreated

	// LifecycleInitialized indicates the component has completed initialization.
	LifecycleInitialized

	// LifecycleRunning indicates the component is actively running.
	LifecycleRunning

	// LifecycleStopping indicates graceful shutdown is in progress.
	LifecycleStopping

	// LifecycleStopped indicates the component has terminated successfully.
	LifecycleStopped

	// LifecycleFailed indicates the component encountered an unrecoverable error.
	LifecycleFailed
)

// String returns the human-readable representation of a lifecycle state.
func (s LifecycleState) String() string {
	switch s {
	case LifecycleCreated:
		return "Created"

	case LifecycleInitialized:
		return "Initialized"

	case LifecycleRunning:
		return "Running"

	case LifecycleStopping:
		return "Stopping"

	case LifecycleStopped:
		return "Stopped"

	case LifecycleFailed:
		return "Failed"

	default:
		return "Unknown"
	}
}

// ============================================================================
// Component Lifecycle Contract
// ============================================================================

// Component defines the lifecycle contract that every Sentinel component
// must implement.
//
// Every future component (Observation, Analysis, Decision, Planner,
// Backend, Telemetry, Reconciliation, etc.) will satisfy this interface.
type Component interface {

	// Name returns the unique name of the component.
	Name() string

	// Init prepares the component and its dependencies.
	Init(ctx context.Context) error

	// Start begins the component's execution.
	Start(ctx context.Context) error

	// Stop gracefully shuts down the component.
	Stop(ctx context.Context) error

	// Health reports whether the component is healthy.
	Health(ctx context.Context) error

	// State returns the current lifecycle state.
	State() LifecycleState
}

// ============================================================================
// Lifecycle Helpers
// ============================================================================

// IsTerminal returns true if the lifecycle state is terminal.
func IsTerminal(state LifecycleState) bool {
	return state == LifecycleStopped ||
		state == LifecycleFailed
}

// IsRunning returns true if the component is actively running.
func IsRunning(state LifecycleState) bool {
	return state == LifecycleRunning
}

package state

import "context"

// Repository defines the contract for Sentinel's
// Shared State Repository (SSR).
//
// The repository is the single source of truth
// for all runtime state.
//
// Implementations must be thread-safe.
type Repository interface {

	// Observation state

	// SetObservation stores the latest observation state.
	SetObservation(ctx context.Context, state ObservationState) error

	// GetObservation retrieves the latest observation state.
	GetObservation(ctx context.Context) (ObservationState, error)

	// Analysis state

	// SetAnalysis stores analysis results.
	SetAnalysis(ctx context.Context, state AnalysisState) error

	// GetAnalysis retrieves analysis results.
	GetAnalysis(ctx context.Context) (AnalysisState, error)

	// Decision state

	// SetDecision stores decision results.
	SetDecision(ctx context.Context, state DecisionState) error

	// GetDecision retrieves decision results.
	GetDecision(ctx context.Context) (DecisionState, error)

	// Planner state

	// SetPlanner stores Adaptive Execution Strategy state.
	SetPlanner(ctx context.Context, state PlannerState) error

	// GetPlanner retrieves planner state.
	GetPlanner(ctx context.Context) (PlannerState, error)

	// Backend state

	// SetBackend stores backend execution state.
	SetBackend(ctx context.Context, state BackendState) error

	// GetBackend retrieves backend execution state.
	GetBackend(ctx context.Context) (BackendState, error)

	// Telemetry state

	// SetTelemetry stores telemetry information.
	SetTelemetry(ctx context.Context, state TelemetryState) error

	// GetTelemetry retrieves telemetry information.
	GetTelemetry(ctx context.Context) (TelemetryState, error)

	// Reconciliation state

	// SetReconciliation stores reconciliation information.
	SetReconciliation(ctx context.Context, state ReconciliationState) error

	// GetReconciliation retrieves reconciliation information.
	GetReconciliation(ctx context.Context) (ReconciliationState, error)

	// Lifecycle

	// Clear removes all stored state.
	Clear(ctx context.Context) error

	// Health verifies repository availability.
	Health(ctx context.Context) error
}

package state

import "context"

// Repository defines the Shared State Repository (SSR).
//
// The SSR is the central in-memory state store shared by all
// Sentinel engines. Each engine owns exactly one section of the
// runtime state and publishes its latest result here.
type Repository interface {

	// -------------------------------------------------------------------------
	// Observation
	// -------------------------------------------------------------------------

	// PublishObservation stores the latest observation snapshot.
	PublishObservation(
		ctx context.Context,
		state *ObservationState,
	) error

	// CurrentObservation returns the latest observation snapshot.
	CurrentObservation(
		ctx context.Context,
	) (*ObservationState, error)

	// -------------------------------------------------------------------------
	// Analysis
	// -------------------------------------------------------------------------

	Analysis(
		ctx context.Context,
	) (AnalysisState, error)

	// -------------------------------------------------------------------------
	// Decision
	// -------------------------------------------------------------------------

	Decision(
		ctx context.Context,
	) (DecisionState, error)

	// -------------------------------------------------------------------------
	// Planner
	// -------------------------------------------------------------------------

	SetPlanner(
		ctx context.Context,
		state PlannerState,
	) error

	Planner(
		ctx context.Context,
	) (PlannerState, error)

	// -------------------------------------------------------------------------
	// Backend
	// -------------------------------------------------------------------------

	SetBackend(
		ctx context.Context,
		state BackendState,
	) error

	Backend(
		ctx context.Context,
	) (BackendState, error)

	// -------------------------------------------------------------------------
	// Telemetry
	// -------------------------------------------------------------------------

	SetTelemetry(
		ctx context.Context,
		state TelemetryState,
	) error

	Telemetry(
		ctx context.Context,
	) (TelemetryState, error)

	// -------------------------------------------------------------------------
	// Reconciliation
	// -------------------------------------------------------------------------

	SetReconciliation(
		ctx context.Context,
		state ReconciliationState,
	) error

	Reconciliation(
		ctx context.Context,
	) (ReconciliationState, error)

	// -------------------------------------------------------------------------
	// Lifecycle
	// -------------------------------------------------------------------------

	Clear(
		ctx context.Context,
	) error

	Health(
		ctx context.Context,
	) error
}

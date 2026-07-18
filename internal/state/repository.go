/* Responsibilities

The Repository interface has exactly four responsibilities:

Publish state into SSR.
Retrieve the latest state from SSR.
Hide the storage implementation.
Remain stable across Sentinel v1, v2, and v3.

It does not:

Manage locks.
Store data.
Perform analysis.
Validate state.

Those belong elsewhere.*/ 

package state

import "context"

// Repository defines the Shared State Repository (SSR).
//
// The SSR is the single source of truth for every runtime state
// produced inside Sentinel.
//
// Implementations must be thread-safe.
type Repository interface {

	// -------------------------------------------------------------------------
	// Observation
	// -------------------------------------------------------------------------

	// PublishObservation atomically publishes a new observation snapshot.
	PublishObservation(
		ctx context.Context,
		state *ObservationState,
	) error

	// CurrentObservation returns the latest published observation snapshot.
	CurrentObservation(
		ctx context.Context,
	) (*ObservationState, error)

	// -------------------------------------------------------------------------
	// Analysis
	// -------------------------------------------------------------------------

	PublishAnalysis(
		ctx context.Context,
		state *AnalysisState,
	) error

	CurrentAnalysis(
		ctx context.Context,
	) (*AnalysisState, error)

	// -------------------------------------------------------------------------
	// Decision
	// -------------------------------------------------------------------------

	PublishDecision(
		ctx context.Context,
		state *DecisionState,
	) error

	CurrentDecision(
		ctx context.Context,
	) (*DecisionState, error)

	// -------------------------------------------------------------------------
	// Planner
	// -------------------------------------------------------------------------

	PublishPlanner(
		ctx context.Context,
		state *PlannerState,
	) error

	CurrentPlanner(
		ctx context.Context,
	) (*PlannerState, error)

	// -------------------------------------------------------------------------
	// Backend
	// -------------------------------------------------------------------------

	PublishBackend(
		ctx context.Context,
		state *BackendState,
	) error

	CurrentBackend(
		ctx context.Context,
	) (*BackendState, error)

	// -------------------------------------------------------------------------
	// Telemetry
	// -------------------------------------------------------------------------

	PublishTelemetry(
		ctx context.Context,
		state *TelemetryState,
	) error

	CurrentTelemetry(
		ctx context.Context,
	) (*TelemetryState, error)

	// -------------------------------------------------------------------------
	// Reconciliation
	// -------------------------------------------------------------------------

	PublishReconciliation(
		ctx context.Context,
		state *ReconciliationState,
	) error

	CurrentReconciliation(
		ctx context.Context,
	) (*ReconciliationState, error)

	// -------------------------------------------------------------------------
	// Store
	// -------------------------------------------------------------------------

	// Clear removes all runtime state.
	Clear(ctx context.Context) error

	// Health reports the health of the repository.
	Health(ctx context.Context) error
}
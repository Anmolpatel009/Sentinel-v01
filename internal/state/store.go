/*
store.go has exactly four responsibilities.

# Store

├── Store runtime state
├── Publish runtime state
├── Serve runtime state
└── Protect concurrent access
*/
package state

import (
	"context"
	"sync"
)

// Store is the in-memory implementation of the
// Sentinel Shared State Repository (SSR).
//
// It owns the latest state produced by each engine
// and provides concurrent, thread-safe access.
type Store struct {
	mu sync.RWMutex

	observation *ObservationState

	analysis       AnalysisState
	decision       DecisionState
	planner        PlannerState
	backend        BackendState
	telemetry      TelemetryState
	reconciliation ReconciliationState
}

// NewStore creates an empty Shared State Repository.
func NewStore() *Store {
	return &Store{}
}

// -----------------------------------------------------------------------------
// Observation
// -----------------------------------------------------------------------------

// PublishObservation stores the latest ObservationState.
func (s *Store) PublishObservation(
	ctx context.Context,
	state *ObservationState,
) error {

	_ = ctx

	s.mu.Lock()
	defer s.mu.Unlock()

	s.observation = state

	return nil
}

// CurrentObservation returns the latest ObservationState.
func (s *Store) CurrentObservation(
	ctx context.Context,
) (*ObservationState, error) {

	_ = ctx

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.observation, nil
}

// -----------------------------------------------------------------------------
// Analysis
// -----------------------------------------------------------------------------

func (s *Store) SetAnalysis(
	ctx context.Context,
	state AnalysisState,
) error {

	_ = ctx

	s.mu.Lock()
	defer s.mu.Unlock()

	s.analysis = state

	return nil
}

func (s *Store) Analysis(
	ctx context.Context,
) (AnalysisState, error) {

	_ = ctx

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.analysis, nil
}

// -----------------------------------------------------------------------------
// Decision
// -----------------------------------------------------------------------------

func (s *Store) SetDecision(
	ctx context.Context,
	state DecisionState,
) error {

	_ = ctx

	s.mu.Lock()
	defer s.mu.Unlock()

	s.decision = state

	return nil
}

func (s *Store) Decision(
	ctx context.Context,
) (DecisionState, error) {

	_ = ctx

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.decision, nil
}

// -----------------------------------------------------------------------------
// Planner
// -----------------------------------------------------------------------------

func (s *Store) SetPlanner(
	ctx context.Context,
	state PlannerState,
) error {

	_ = ctx

	s.mu.Lock()
	defer s.mu.Unlock()

	s.planner = state

	return nil
}

func (s *Store) Planner(
	ctx context.Context,
) (PlannerState, error) {

	_ = ctx

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.planner, nil
}

// -----------------------------------------------------------------------------
// Backend
// -----------------------------------------------------------------------------

func (s *Store) SetBackend(
	ctx context.Context,
	state BackendState,
) error {

	_ = ctx

	s.mu.Lock()
	defer s.mu.Unlock()

	s.backend = state

	return nil
}

func (s *Store) Backend(
	ctx context.Context,
) (BackendState, error) {

	_ = ctx

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.backend, nil
}

// -----------------------------------------------------------------------------
// Telemetry
// -----------------------------------------------------------------------------

func (s *Store) SetTelemetry(
	ctx context.Context,
	state TelemetryState,
) error {

	_ = ctx

	s.mu.Lock()
	defer s.mu.Unlock()

	s.telemetry = state

	return nil
}

func (s *Store) Telemetry(
	ctx context.Context,
) (TelemetryState, error) {

	_ = ctx

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.telemetry, nil
}

// -----------------------------------------------------------------------------
// Reconciliation
// -----------------------------------------------------------------------------

func (s *Store) SetReconciliation(
	ctx context.Context,
	state ReconciliationState,
) error {

	_ = ctx

	s.mu.Lock()
	defer s.mu.Unlock()

	s.reconciliation = state

	return nil
}

func (s *Store) Reconciliation(
	ctx context.Context,
) (ReconciliationState, error) {

	_ = ctx

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.reconciliation, nil
}

// -----------------------------------------------------------------------------
// Lifecycle
// -----------------------------------------------------------------------------

// Clear removes all runtime state.
func (s *Store) Clear(
	ctx context.Context,
) error {

	_ = ctx

	s.mu.Lock()
	defer s.mu.Unlock()

	s.observation = nil
	s.analysis = AnalysisState{}
	s.decision = DecisionState{}
	s.planner = PlannerState{}
	s.backend = BackendState{}
	s.telemetry = TelemetryState{}
	s.reconciliation = ReconciliationState{}

	return nil
}

// Health reports the health of the repository.
//
// Sentinel v1 uses an in-memory repository, so there are
// no external dependencies to verify.
func (s *Store) Health(
	ctx context.Context,
) error {

	_ = ctx

	return nil
}

// Compile-time verification.
var _ Repository = (*Store)(nil)

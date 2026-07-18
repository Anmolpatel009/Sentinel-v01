package state

import (
	"context"
	"sync"
)

// Store is the in-memory implementation
// of the Sentinel Shared State Repository.
//
// It is thread-safe and uses RWMutex
// to protect concurrent access.
type Store struct {
	mu sync.RWMutex

	observation ObservationState
	analysis    AnalysisState
	decision    DecisionState
	planner     PlannerState
	backend     BackendState
	telemetry   TelemetryState
	reconcile   ReconciliationState
}

// NewStore creates a new empty SSR instance.
func NewStore() *Store {
	return &Store{}
}

// -------------------------
// Observation
// -------------------------

func (s *Store) SetObservation(
	ctx context.Context,
	state ObservationState,
) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.observation = state

	return nil
}

func (s *Store) GetObservation(
	ctx context.Context,
) (ObservationState, error) {

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.observation, nil
}

// -------------------------
// Analysis
// -------------------------

func (s *Store) SetAnalysis(
	ctx context.Context,
	state AnalysisState,
) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.analysis = state

	return nil
}

func (s *Store) GetAnalysis(
	ctx context.Context,
) (AnalysisState, error) {

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.analysis, nil
}

// -------------------------
// Decision
// -------------------------

func (s *Store) SetDecision(
	ctx context.Context,
	state DecisionState,
) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.decision = state

	return nil
}

func (s *Store) GetDecision(
	ctx context.Context,
) (DecisionState, error) {

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.decision, nil
}

// -------------------------
// Planner
// -------------------------

func (s *Store) SetPlanner(
	ctx context.Context,
	state PlannerState,
) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.planner = state

	return nil
}

func (s *Store) GetPlanner(
	ctx context.Context,
) (PlannerState, error) {

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.planner, nil
}

// -------------------------
// Backend
// -------------------------

func (s *Store) SetBackend(
	ctx context.Context,
	state BackendState,
) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.backend = state

	return nil
}

func (s *Store) GetBackend(
	ctx context.Context,
) (BackendState, error) {

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.backend, nil
}

// -------------------------
// Telemetry
// -------------------------

func (s *Store) SetTelemetry(
	ctx context.Context,
	state TelemetryState,
) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.telemetry = state

	return nil
}

func (s *Store) GetTelemetry(
	ctx context.Context,
) (TelemetryState, error) {

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.telemetry, nil
}

// -------------------------
// Reconciliation
// -------------------------

func (s *Store) SetReconciliation(
	ctx context.Context,
	state ReconciliationState,
) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.reconcile = state

	return nil
}

func (s *Store) GetReconciliation(
	ctx context.Context,
) (ReconciliationState, error) {

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.reconcile, nil
}

// -------------------------
// Lifecycle
// -------------------------

// Clear removes all stored state.
func (s *Store) Clear(
	ctx context.Context,
) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.observation = ObservationState{}
	s.analysis = AnalysisState{}
	s.decision = DecisionState{}
	s.planner = PlannerState{}
	s.backend = BackendState{}
	s.telemetry = TelemetryState{}
	s.reconcile = ReconciliationState{}

	return nil
}

// Health verifies SSR availability.
func (s *Store) Health(
	ctx context.Context,
) error {

	// v1:
	// Memory store has no external dependency.

	return nil
}

// Compile-time interface verification.
var _ Repository = (*Store)(nil)

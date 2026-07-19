/*
	             Manager
	                │
	   Allocate ObservationState
	                │
	┌───────────────┼────────────────┐
	▼               ▼                ▼

CPU           Memory          Network

	▼               ▼                ▼

Process      Filesystem      (future...)

	               │
	               ▼
	Repository.PublishObservation()
*/
package observation

import (
	"context"

	"github.com/Anmolpatel009/Sentinel-v01/internal/observation/cpu"
	"github.com/Anmolpatel009/Sentinel-v01/internal/observation/filesystem"
	"github.com/Anmolpatel009/Sentinel-v01/internal/observation/memory"
	"github.com/Anmolpatel009/Sentinel-v01/internal/observation/network"
	"github.com/Anmolpatel009/Sentinel-v01/internal/observation/process"
	"github.com/Anmolpatel009/Sentinel-v01/internal/state"
)

// Manager coordinates the Observation Engine.
//
// The Manager owns all built-in collectors, creates exactly one
// ObservationState per observation cycle, invokes every collector,
// and publishes the completed snapshot into the Shared State
// Repository (SSR).
//
// Sentinel v1 intentionally owns the collectors directly because
// there is only one implementation of each collector. Runtime
// injection and collector interfaces are deferred until Sentinel v2.
type Manager struct {
	repository state.Repository

	cpu        *cpu.Collector
	memory     *memory.Collector
	network    *network.Collector
	process    *process.Collector
	filesystem *filesystem.Collector
}

// NewManager creates a new Observation Manager.
func NewManager(
	repository state.Repository,
) *Manager {

	return &Manager{
		repository: repository,

		cpu:        cpu.NewCollector(),
		memory:     memory.NewCollector(),
		network:    network.NewCollector(),
		process:    process.NewCollector(),
		filesystem: filesystem.NewCollector(),
	}
}

// Collect executes one complete observation cycle.
func (m *Manager) Collect(
	ctx context.Context,
) error {

	// Allocate exactly one snapshot for this observation cycle.
	snapshot := &state.ObservationState{}

	// CPU
	if err := m.cpu.Collect(ctx, snapshot); err != nil {
		return err
	}

	// Memory
	if err := m.memory.Collect(ctx, snapshot); err != nil {
		return err
	}

	// Network
	if err := m.network.Collect(ctx, snapshot); err != nil {
		return err
	}

	// Process
	if err := m.process.Collect(ctx, snapshot); err != nil {
		return err
	}

	// Filesystem
	if err := m.filesystem.Collect(ctx, snapshot); err != nil {
		return err
	}

	// Publish the completed snapshot atomically.
	return m.repository.PublishObservation(ctx, snapshot)
}

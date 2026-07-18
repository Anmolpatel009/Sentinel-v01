/*

Runtime
    │
    ▼
Observation Component
    │
    ▼
Observation Manager
    │
    ├──────── CPU Collector
    ├──────── Memory Collector
    ├──────── Network Collector
    ├──────── Process Collector
    ├──────── Filesystem Collector
    │
    ▼
Publisher
    │
    ▼
SSR
*/
package observation

import (
	"context"

	"github.com/Anmolpatel009/Sentinel-v01/internal/observation/cpu"
	"github.com/Anmolpatel009/Sentinel-v01/internal/observation/filesystem"
	"github.com/Anmolpatel009/Sentinel-v01/internal/observation/memory"
	"github.com/Anmolpatel009/Sentinel-v01/internal/observation/network"
	"github.com/Anmolpatel009/Sentinel-v01/internal/observation/process"
)

// Manager coordinates one complete observation cycle.
type Manager struct {
	cpu        *cpu.Collector
	memory     *memory.Collector
	network    *network.Collector
	process    *process.Collector
	filesystem *filesystem.Collector

	publisher *Publisher
}

// NewManager creates a new Observation Manager.
func NewManager(
	cpuCollector *cpu.Collector,
	memoryCollector *memory.Collector,
	networkCollector *network.Collector,
	processCollector *process.Collector,
	filesystemCollector *filesystem.Collector,
	publisher *Publisher,
) *Manager {
	return &Manager{
		cpu:        cpuCollector,
		memory:     memoryCollector,
		network:    networkCollector,
		process:    processCollector,
		filesystem: filesystemCollector,
		publisher:  publisher,
	}
}

// Observe performs one complete observation cycle.
func (m *Manager) Observe(ctx context.Context) error {

	snapshot := ObservationSnapshot{}

	if err := m.cpu.Collect(ctx, &snapshot); err != nil {
		return err
	}

	if err := m.memory.Collect(ctx, &snapshot); err != nil {
		return err
	}

	if err := m.network.Collect(ctx, &snapshot); err != nil {
		return err
	}

	if err := m.process.Collect(ctx, &snapshot); err != nil {
		return err
	}

	if err := m.filesystem.Collect(ctx, &snapshot); err != nil {
		return err
	}

	return m.publisher.Publish(ctx, snapshot)
}
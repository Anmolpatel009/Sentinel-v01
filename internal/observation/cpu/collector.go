package cpu

import (
	"context"

	"github.com/Anmolpatel009/Sentinel-v01/internal/state"
)

// Collector gathers CPU metrics from the operating system.
//
// It owns only the CPU portion of ObservationState.
type Collector struct{}

// NewCollector creates a new CPU collector.
func NewCollector() *Collector {
	return &Collector{}
}

// Collect gathers CPU information and populates the CPU section
// of the observation snapshot.
func (c *Collector) Collect(
	ctx context.Context,
	snapshot *state.ObservationState,
) error {

	cpuState, err := c.collect(ctx)
	if err != nil {
		return err
	}

	snapshot.CPU = cpuState

	return nil
}

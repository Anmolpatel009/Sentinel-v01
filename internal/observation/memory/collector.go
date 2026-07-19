package memory

import (
	"context"

	"github.com/Anmolpatel009/Sentinel-v01/internal/state"
)

// Collector gathers memory metrics from the operating system.
//
// It owns only the Memory portion of ObservationState.
type Collector struct{}

// NewCollector creates a new Memory collector.
func NewCollector() *Collector {
	return &Collector{}
}

// Collect gathers memory information and populates the memory
// section of the observation snapshot.
func (c *Collector) Collect(
	ctx context.Context,
	snapshot *state.ObservationState,
) error {

	memoryState, err := c.collect(ctx)
	if err != nil {
		return err
	}

	snapshot.Memory = memoryState

	return nil
}

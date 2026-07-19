package process

import (
	"context"

	"github.com/Anmolpatel009/Sentinel-v01/internal/state"
)

// Collector gathers process metrics from the operating system.
//
// It owns only the Process portion of ObservationState.
type Collector struct{}

// NewCollector creates a new Process collector.
func NewCollector() *Collector {
	return &Collector{}
}

// Collect gathers process information and populates the process
// section of the observation snapshot.
func (c *Collector) Collect(
	ctx context.Context,
	snapshot *state.ObservationState,
) error {

	processState, err := c.collect(ctx)
	if err != nil {
		return err
	}

	snapshot.Process = processState

	return nil
}

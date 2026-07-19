package filesystem

import (
	"context"

	"github.com/Anmolpatel009/Sentinel-v01/internal/state"
)

// Collector gathers filesystem metrics from the operating system.
//
// It owns only the Filesystem portion of ObservationState.
type Collector struct{}

// NewCollector creates a new Filesystem collector.
func NewCollector() *Collector {
	return &Collector{}
}

// Collect gathers filesystem information and populates the
// filesystem section of the observation snapshot.
func (c *Collector) Collect(
	ctx context.Context,
	snapshot *state.ObservationState,
) error {

	filesystemState, err := c.collect(ctx)
	if err != nil {
		return err
	}

	snapshot.Filesystem = filesystemState

	return nil
}

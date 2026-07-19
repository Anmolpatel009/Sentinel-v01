package network

import (
	"context"

	"github.com/Anmolpatel009/Sentinel-v01/internal/state"
)

// Collector gathers network metrics from the operating system.
//
// It owns only the Network portion of ObservationState.
type Collector struct{}

// NewCollector creates a new Network collector.
func NewCollector() *Collector {
	return &Collector{}
}

// Collect gathers network information and populates the network
// section of the observation snapshot.
func (c *Collector) Collect(
	ctx context.Context,
	snapshot *state.ObservationState,
) error {

	networkState, err := c.collect(ctx)
	if err != nil {
		return err
	}

	snapshot.Network = networkState

	return nil
}

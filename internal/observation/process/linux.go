//go:build linux

package process

import (
	"context"

	"github.com/Anmolpatel009/Sentinel-v01/internal/state"
)

// collect gathers process information from Linux.
//
// v1 implementation returns a placeholder state.
//
// Future versions will collect metrics from:
//
//   - /proc
//   - /proc/stat
//
// without changing the public collector API.
func (c *Collector) collect(
	ctx context.Context,
) (state.ProcessState, error) {

	return state.ProcessState{}, nil
}

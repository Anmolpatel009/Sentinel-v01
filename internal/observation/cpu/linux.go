//go:build linux

package cpu

import (
	"context"

	"github.com/Anmolpatel009/Sentinel-v01/internal/state"
)

// collect gathers CPU information from Linux.
//
// v1 implementation returns a placeholder state.
// Future versions will parse:
//
//   - /proc/stat
//   - /proc/loadavg
//
// without changing the public collector API.
func (c *Collector) collect(
	ctx context.Context,
) (state.CPUState, error) {

	return state.CPUState{}, nil
}

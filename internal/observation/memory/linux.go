//go:build linux

package memory

import (
	"context"

	"github.com/Anmolpatel009/Sentinel-v01/internal/state"
)

// collect gathers memory information from Linux.
//
// v1 implementation returns a placeholder state.
// Future versions will parse:
//
//   - /proc/meminfo
func (c *Collector) collect(
	ctx context.Context,
) (state.MemoryState, error) {

	return state.MemoryState{}, nil
}

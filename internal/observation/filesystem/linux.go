//go:build linux

package filesystem

import (
	"context"

	"github.com/Anmolpatel009/Sentinel-v01/internal/state"
)

// collect gathers filesystem information from Linux.
//
// v1 implementation returns a placeholder state.
//
// Future versions will collect metrics using:
//
//   - statfs(2)
//   - /proc/self/mountinfo
//   - /proc/mounts
//
// without changing the public collector API.
func (c *Collector) collect(
	ctx context.Context,
) (state.FilesystemState, error) {

	return state.FilesystemState{}, nil
}

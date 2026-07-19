//go:build linux

package network

import (
	"context"

	"github.com/Anmolpatel009/Sentinel-v01/internal/state"
)

// collect gathers network information from Linux.
//
// v1 implementation returns a placeholder state.
//
// Future versions will collect metrics from:
//
//   - /proc/net/dev
//   - /proc/net/snmp
//   - /proc/net/netstat
//
// Sentinel v2 may combine these with eBPF.
//
// Sentinel v3 will primarily use eBPF/XDP while
// preserving the same public collector API.
func (c *Collector) collect(
	ctx context.Context,
) (state.NetworkState, error) {

	return state.NetworkState{}, nil
}

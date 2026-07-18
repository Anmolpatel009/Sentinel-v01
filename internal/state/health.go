package state

import "context"

// Healthy reports whether the Shared State Repository (SSR)
// is currently healthy.
//
// It is a convenience wrapper around Health() and is intended
// for quick runtime health checks.
//
// In Sentinel v1, the SSR is an in-memory repository with no
// external dependencies. Therefore, a nil error from Health()
// indicates the subsystem is healthy.
//
// Future versions may extend Health() to verify:
//
//   - Memory usage
//   - Lock contention
//   - Snapshot subsystem
//   - Persistent storage
//   - Replication status
//   - Distributed state synchronization
func (c *Component) Healthy(ctx context.Context) bool {
	return c.Health(ctx) == nil
}

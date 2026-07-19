package observation

import "context"

// Healthy reports whether the Observation Engine
// is currently healthy.
func (c *Component) Healthy(
	ctx context.Context,
) bool {

	return c.Health(ctx) == nil
}

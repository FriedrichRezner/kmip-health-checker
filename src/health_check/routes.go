package health_check

import (
	"github.com/friedrichrezner/kmip-health-checker/src/health_check/interfaces"

	"flamingo.me/flamingo/v3/framework/web"
)

// routes implements the Flamingo routes interface and defines the routes for the health check module
type routes struct {
	healthCheckController *interfaces.HealthCheckController
}

// Inject dependencies
func (r *routes) Inject(c *interfaces.HealthCheckController) *routes {
	r.healthCheckController = c

	return r
}

// Routes maps the path to the specific controller action
func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/health", "health")
	registry.HandleGet("health", r.healthCheckController.HealthCheck)
}

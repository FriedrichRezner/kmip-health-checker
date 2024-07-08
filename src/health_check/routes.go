package health_check

import (
	"flamingo.me/flamingo/v3/framework/web"
	"kmip-health-checker/src/health_check/interfaces"
)

type routes struct {
	healthCheckController *interfaces.HealthCheckController
}

func (r *routes) Inject(c *interfaces.HealthCheckController) *routes {
	r.healthCheckController = c

	return r
}

func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/health", "health")
	registry.HandleGet("health", r.healthCheckController.HealthCheck)
}

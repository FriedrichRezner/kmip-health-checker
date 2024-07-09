package health_check

import (
	"github.com/friedrichrezner/kmip-health-checker/src/health_check/application"
	"github.com/friedrichrezner/kmip-health-checker/src/health_check/domain"
	"github.com/friedrichrezner/kmip-health-checker/src/health_check/infrastructure"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
)

// Module for health_check
type Module struct{}

// Configure defines the dependencies needed for the health check module
func (*Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))

	injector.Bind(new(infrastructure.KMIPClient)).To(infrastructure.KMIPClientImpl{})
	injector.Bind(new(domain.KMIPRepository)).To(infrastructure.KMIPAdapter{})
	injector.Bind(new(domain.HealthChecker)).To(application.HealthCheckService{})
}

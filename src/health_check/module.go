package health_check

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	"kmip-health-checker/src/health_check/application"
	"kmip-health-checker/src/health_check/domain"
	"kmip-health-checker/src/health_check/infrastructure"
)

type Module struct{}

func (*Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))

	injector.Bind(new(infrastructure.KMIPClient)).To(infrastructure.KMIPClientImpl{})
	injector.Bind(new(domain.KMIPRepository)).To(infrastructure.KMIPAdapter{})
	injector.Bind(new(domain.HealthChecker)).To(application.HealthCheckService{})
}

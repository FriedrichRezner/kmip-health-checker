package interfaces

import (
	"context"
	"errors"
	"flamingo.me/flamingo/v3/framework/web"
	"kmip-health-checker/src/health_check/domain"
	"strconv"
)

const (
	defaultAmount  = 1
	amountQueryKey = "amount"
)

type (
	HealthCheckResponse struct {
		CreateDuration  float64 `json:"create_duration"`
		DestroyDuration float64 `json:"destroy_duration"`
		Amount          int     `json:"amount"`
	}
	HealthCheckController struct {
		responder     *web.Responder
		healthChecker domain.HealthChecker
	}
)

func (c *HealthCheckController) Inject(r *web.Responder, hc domain.HealthChecker) {
	c.responder = r
	c.healthChecker = hc
}

func (c *HealthCheckController) HealthCheck(ctx context.Context, r *web.Request) web.Result {
	amount := defaultAmount

	queryAmount, err := r.Query1(amountQueryKey)
	if err == nil {
		amountInt, err := strconv.Atoi(queryAmount)
		if err == nil {
			amount = amountInt
		}
	}

	healthCheckResult, err := c.healthChecker.PerformCheck(ctx, amount)
	if err != nil {
		return c.responder.ServerError(errors.New("health check failed. see logs for more information"))
	}

	return c.responder.Data(mapResponse(*healthCheckResult))
}

func mapResponse(healthCheckResult domain.HealthCheckResult) *HealthCheckResponse {
	return &HealthCheckResponse{
		CreateDuration:  healthCheckResult.CreateDuration.Seconds(),
		DestroyDuration: healthCheckResult.DestroyDuration.Seconds(),
		Amount:          healthCheckResult.Amount,
	}
}

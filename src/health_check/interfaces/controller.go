package interfaces

import (
	"context"
	"errors"
	"strconv"

	"github.com/friedrichrezner/kmip-health-checker/src/health_check/domain"

	"flamingo.me/flamingo/v3/framework/web"
)

const (
	defaultAmount  = 1
	amountQueryKey = "amount"
)

type (
	// HealthCheckResponse is the response struct for the rest endpoint providing float64 instead of time.Duration
	HealthCheckResponse struct {
		CreateDuration  float64 `json:"create_duration"`
		DestroyDuration float64 `json:"destroy_duration"`
		Amount          int     `json:"amount"`
	}

	// HealthCheckController is the controller for the health check endpoint
	// It only uses the HealthChecker to perform the health check
	HealthCheckController struct {
		responder     *web.Responder
		healthChecker domain.HealthChecker
	}
)

// Inject dependencies
func (c *HealthCheckController) Inject(r *web.Responder, hc domain.HealthChecker) {
	c.responder = r
	c.healthChecker = hc
}

// HealthCheck is the controller method for the health check endpoint
// It checks if there is a query parameter "amount" and otherwise uses the configured default amount
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

// mapResponse maps the domain.HealthCheckResult to the HealthCheckResponse
func mapResponse(healthCheckResult domain.HealthCheckResult) *HealthCheckResponse {
	return &HealthCheckResponse{
		CreateDuration:  healthCheckResult.CreateDuration.Seconds(),
		DestroyDuration: healthCheckResult.DestroyDuration.Seconds(),
		Amount:          healthCheckResult.Amount,
	}
}

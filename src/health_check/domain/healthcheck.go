package domain

import (
	"context"
	"time"
)

type (
	// HealthCheckResult stores the duration that was needed to create and destroy the keys
	HealthCheckResult struct {
		CreateDuration  time.Duration
		DestroyDuration time.Duration
		Amount          int
	}
	// HealthChecker is the interface for performing an amount of checks and returning the result
	HealthChecker interface {
		PerformCheck(ctx context.Context, amount int) (*HealthCheckResult, error)
	}
)

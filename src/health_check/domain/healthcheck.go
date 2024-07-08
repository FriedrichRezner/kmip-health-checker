package domain

import (
	"context"
	"time"
)

type (
	HealthCheckResult struct {
		CreateDuration  time.Duration
		DestroyDuration time.Duration
		Amount          int
	}
	HealthChecker interface {
		PerformCheck(ctx context.Context, amount int) (*HealthCheckResult, error)
	}
)

package application

import (
	"context"
	"fmt"
	"time"

	"github.com/friedrichrezner/kmip-health-checker/src/health_check/domain"

	"flamingo.me/flamingo/v3/framework/flamingo"
)

// HealthCheckService is the application service for the health check
type HealthCheckService struct {
	kmipRepo domain.KMIPRepository
	logger   flamingo.Logger
}

// Inject dependencies
func (h *HealthCheckService) Inject(kmipRepo domain.KMIPRepository, l flamingo.Logger) {
	h.kmipRepo = kmipRepo
	h.logger = l
}

// PerformCheck performs the health check based on the amount of keys to create
// First the keys are created and stored in a string slice
// After logging the needed time the keys are deleted
func (h *HealthCheckService) PerformCheck(ctx context.Context, amount int) (*domain.HealthCheckResult, error) {
	h.logger.WithContext(ctx).Debugf("start creating %d keys", amount)

	result := domain.HealthCheckResult{Amount: amount}

	timer := time.Now()
	createdIds, err := h.create(ctx, amount)
	if err != nil {
		h.logger.WithContext(ctx).Error(fmt.Sprintf("error creating keys: %v", err))
		return nil, err
	}

	result.CreateDuration = time.Since(timer)
	timer = time.Now()

	err = h.delete(ctx, createdIds)
	if err != nil {
		h.logger.WithContext(ctx).Error(fmt.Sprintf("error deleting keys: %v", err))
		return nil, err
	}

	result.DestroyDuration = time.Since(timer)

	return &result, nil
}

// create creates the amount of keys and returns the ids
func (h *HealthCheckService) create(ctx context.Context, amount int) ([]string, error) {
	var createdIds []string

	for i := 0; i < amount; i++ {
		id, err := h.kmipRepo.Create(ctx)
		if err != nil {
			return nil, err
		}

		h.logger.WithContext(ctx).Debugf("Created key with id: %s", id)

		createdIds = append(createdIds, id)
	}

	return createdIds, nil
}

// delete deletes the keys with the given ids
func (h *HealthCheckService) delete(ctx context.Context, ids []string) error {
	for _, id := range ids {
		err := h.kmipRepo.Destroy(ctx, id)
		if err != nil {
			return err
		}
		h.logger.WithContext(ctx).Debugf("Destroyed key with id: %s", id)
	}

	return nil
}

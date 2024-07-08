package application

import (
	"context"
	"flamingo.me/flamingo/v3/framework/flamingo"
	"fmt"
	"kmip-health-checker/src/health_check/domain"
	"time"
)

type HealthCheckService struct {
	kmipRepo domain.KMIPRepository
	logger   flamingo.Logger
}

func (h *HealthCheckService) Inject(kmipRepo domain.KMIPRepository, l flamingo.Logger) {
	h.kmipRepo = kmipRepo
	h.logger = l
}

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

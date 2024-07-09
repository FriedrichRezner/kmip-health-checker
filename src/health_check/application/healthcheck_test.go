package application

import (
	"context"
	"errors"
	"testing"

	"flamingo.me/flamingo/v3/framework/flamingo"
	"github.com/friedrichrezner/kmip-health-checker/src/health_check/domain"
	"github.com/stretchr/testify/suite"
)

type HealthCheckServiceTestSuite struct {
	suite.Suite

	context            context.Context
	kmipRepoMock       *domain.MockKMIPRepository
	healthCheckService *HealthCheckService
}

func (t *HealthCheckServiceTestSuite) SetupTest() {
	t.context = context.Background()
	t.kmipRepoMock = &domain.MockKMIPRepository{}
	t.healthCheckService = &HealthCheckService{}

	t.healthCheckService.Inject(t.kmipRepoMock, flamingo.NullLogger{})
}

func TestHealthCheckServiceTestSuite(t *testing.T) {
	suite.Run(t, new(HealthCheckServiceTestSuite))
}

func (t *HealthCheckServiceTestSuite) TestPerformCheck() {
	t.Run("success", func() {
		expectedResult := domain.HealthCheckResult{
			Amount: 2,
		}

		t.kmipRepoMock.On("Create", t.context).Return("1", nil).Once()
		t.kmipRepoMock.On("Create", t.context).Return("2", nil).Once()
		t.kmipRepoMock.On("Destroy", t.context, "1").Return(nil).Once()
		t.kmipRepoMock.On("Destroy", t.context, "2").Return(nil).Once()

		res, err := t.healthCheckService.PerformCheck(t.context, 2)
		t.NoError(err)
		t.Equal(res.Amount, expectedResult.Amount)
	})

	t.Run("creating key fails", func() {
		t.kmipRepoMock.On("Create", t.context).Return("", errors.New("oh no")).Once()

		res, err := t.healthCheckService.PerformCheck(t.context, 2)
		t.Error(err)
		t.Empty(res)
	})

	t.Run("destroying key fails", func() {
		t.kmipRepoMock.On("Create", t.context).Return("1", nil).Once()
		t.kmipRepoMock.On("Destroy", t.context, "1").Return(errors.New("error")).Once()

		res, err := t.healthCheckService.PerformCheck(t.context, 1)
		t.Error(err)
		t.Empty(res)
	})
}

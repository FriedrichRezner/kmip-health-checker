package interfaces

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/friedrichrezner/kmip-health-checker/src/health_check/domain"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/stretchr/testify/suite"
)

type HealthCheckControllerTestSuite struct {
	suite.Suite

	context               context.Context
	healthCheckerMock     *domain.MockHealthChecker
	healthCheckController *HealthCheckController
}

func (t *HealthCheckControllerTestSuite) SetupTest() {
	t.context = context.Background()
	t.healthCheckerMock = &domain.MockHealthChecker{}
	t.healthCheckController = &HealthCheckController{}

	t.healthCheckController.Inject(&web.Responder{}, t.healthCheckerMock)
}

func TestHealthCheckControllerTestSuite(t *testing.T) {
	suite.Run(t, new(HealthCheckControllerTestSuite))
}

func (t *HealthCheckControllerTestSuite) TestHealthCheck() {
	t.Run("success with default amount", func() {
		expected := domain.HealthCheckResult{Amount: 1}
		httpReq, _ := http.NewRequest(http.MethodGet, "", nil)
		req := web.CreateRequest(
			httpReq,
			web.EmptySession(),
		)

		t.healthCheckerMock.On("PerformCheck", t.context, defaultAmount).Return(&expected, nil).Once()

		res := t.healthCheckController.HealthCheck(t.context, req)

		_, ok := res.(web.Result)

		t.True(ok)
	})

	t.Run("success with specified amount", func() {
		expected := domain.HealthCheckResult{Amount: 10}
		httpReq, _ := http.NewRequest(http.MethodGet, "localhost?amount=10", nil)
		req := web.CreateRequest(
			httpReq,
			web.EmptySession(),
		)

		t.healthCheckerMock.On("PerformCheck", t.context, 10).Return(&expected, nil).Once()

		res := t.healthCheckController.HealthCheck(t.context, req)

		_, ok := res.(web.Result)

		t.True(ok)
	})

	t.Run("health check fails", func() {
		httpReq, _ := http.NewRequest(http.MethodGet, "", nil)
		req := web.CreateRequest(
			httpReq,
			web.EmptySession(),
		)

		t.healthCheckerMock.On("PerformCheck", t.context, defaultAmount).Return(nil, errors.New("error")).Once()

		res := t.healthCheckController.HealthCheck(t.context, req)

		_, ok := res.(web.Result)

		t.True(ok)
	})
}

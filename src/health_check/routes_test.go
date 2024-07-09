package health_check

import (
	"testing"

	"github.com/friedrichrezner/kmip-health-checker/src/health_check/interfaces"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	r := routes{}
	r.Inject(&interfaces.HealthCheckController{})

	reg := web.NewRegistry()
	r.Routes(reg)

	handler := reg.GetRoutes()

	assert.Len(t, handler, 1)
}

package health_check

import (
	"testing"

	"flamingo.me/flamingo/v3/framework/config"
	"github.com/stretchr/testify/assert"
)

func TestModule(t *testing.T) {
	err := config.TryModules(config.Map{}, new(Module))
	assert.NoError(t, err)
}

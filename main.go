package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/requestlogger"
	"kmip-health-checker/src/health_check"
)

func main() {
	flamingo.App([]dingo.Module{
		new(requestlogger.Module),
		new(health_check.Module),
	})
}

package main

import (
	"medical-appointments/cmd/api"
	"medical-appointments/internal/infrastructure/config"
)

func main() {
	// Initialize dependencies
	deps, err := config.InitializeDependencies()
	if err != nil {
		deps.Logger.Error(err.Error())
		return
	}

	// Start api
	err = api.InitApplication(deps)
	if err != nil {
		deps.Logger.Error(err.Error())
		return
	}
}

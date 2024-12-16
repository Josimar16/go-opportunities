package main

import (
	"github.com/Josimar16/go-opportunities/config"
	"github.com/Josimar16/go-opportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize the configuration
	err := config.Init()

	if err != nil {
		logger.Errorf("Error initializing the configuration: %v", err)
		return
	}

	// Initialize the router
	router.Initialize()
}

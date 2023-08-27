package main

import (
	"github.com/steixeira93/gopportunities/config"
	"github.com/steixeira93/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	// Initialize Congifs
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
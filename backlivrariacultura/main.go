package main

import (
	"github.com/rafabcanedo/LibraryCulture/config"
	"github.com/rafabcanedo/LibraryCulture/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Config
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	router.Initialize()
}

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
	error := config.Init()
	if error != nil {
		logger.Errorf("Config initialization error: %v", error)
		return
	}

	router.Initialize()
}

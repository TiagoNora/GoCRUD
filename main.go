package main

import (
	"github.com/TiagoNora/GoCRUD/config"
	"github.com/TiagoNora/GoCRUD/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.NewLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}

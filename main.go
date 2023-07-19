package main

import (
	"github.com/TiagoNora/GoCRUD/config"
	"github.com/TiagoNora/GoCRUD/router"
)

var (
	logger *config.Logger
)

// @title           Swagger GO API
// @version         1.0
// @description     This is a example GO API.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey bearerToken
// @in header
// @name Authorization
func main() {
	logger = config.NewLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}

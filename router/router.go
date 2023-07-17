package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()
	initializeRoutes(router)
	router.Run("127.0.0.1:8080")
}

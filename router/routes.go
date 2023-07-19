package router

import (
	"github.com/TiagoNora/GoCRUD/docs"
	"github.com/TiagoNora/GoCRUD/handler"
	"github.com/TiagoNora/GoCRUD/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group("/api/v1")
	{
		v1.GET("/products", handler.ListProducts)
		v1.GET("/product", handler.ShowProduct)
		v1.POST("/user", handler.CreateUser)
		v1.POST("/user/token", handler.GenerateToken)
	}
	v2 := router.Group("/api/v1/secured").Use(middleware.Auth())
	{
		v2.POST("/product", handler.CreateProduct)
		v2.DELETE("/product", handler.DeleteProduct)
		v2.PUT("/product", handler.UpdateProduct)
		v2.GET("/users", handler.ListUsers)
		v2.GET("/user", handler.ShowUser)
		v2.DELETE("/user", handler.DeleteUser)
		v2.PUT("/user", handler.UpdateUser)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

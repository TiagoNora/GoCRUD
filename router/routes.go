package router

import (
	"github.com/TiagoNora/GoCRUD/docs"
	"github.com/TiagoNora/GoCRUD/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group("/api/v1")
	{
		v1.GET("/products", handler.ListProducts)
		v1.GET("/product", handler.ShowProduct)
		v1.POST("/product", handler.CreateProduct)
		v1.DELETE("/product", handler.DeleteProduct)
		v1.PUT("/product", handler.UpdateProduct)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

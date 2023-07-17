package router

import (
	"github.com/TiagoNora/GoCRUD/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/products", handler.ListProducts)
		v1.GET("/product", handler.ShowProduct)
		v1.POST("/products", handler.CreateProduct)
		v1.DELETE("/products", handler.DeleteProduct)
		v1.PUT("/products", handler.UpdateProduct)
	}
}

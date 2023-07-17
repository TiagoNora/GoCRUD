package router

import (
	"github.com/TiagoNora/GoCRUD/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutesProduct(v1 *gin.RouterGroup) {
	v1.GET("/products", handler.ListProducts)
	v1.GET("/product", handler.ShowProduct)
	v1.POST("/product", handler.CreateProduct)
	v1.DELETE("/product", handler.DeleteProduct)
	v1.PUT("/product", handler.UpdateProduct)
}

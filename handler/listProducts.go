package handler

import (
	"net/http"

	"github.com/TiagoNora/GoCRUD/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List products
// @Description List all products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {object} ListProductsResponse
// @Failure 500 {object} ErrorResponse
// @Router /products [get]
func ListProducts(ctx *gin.Context) {
	products := []schemas.Product{}

	if err := db.Find(&products).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing products")
		return
	}

	sendSuccess(ctx, "list-products", products)
}

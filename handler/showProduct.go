package handler

import (
	"net/http"

	"github.com/TiagoNora/GoCRUD/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show product
// @Description Show a product
// @Tags Products
// @Accept json
// @Produce json
// @Param id query string true "Product identification"
// @Success 200 {object} ShowProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /product [get]
func ShowProduct(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	product := schemas.Product{}
	if err := db.First(&product, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "product not found")
		return
	}

	sendSuccess(ctx, "show-product", product)
}

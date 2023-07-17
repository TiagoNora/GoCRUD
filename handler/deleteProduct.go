package handler

import (
	"fmt"
	"net/http"

	"github.com/TiagoNora/GoCRUD/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteProduct(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	product := schemas.Product{}
	// Find Opening
	if err := db.First(&product, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("product with id: %s not found", id))
		return
	}
	// Delete Opening
	if err := db.Delete(&product).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting product with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-product", product)
}

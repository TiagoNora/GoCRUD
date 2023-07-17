package handler

import (
	"net/http"

	"github.com/TiagoNora/GoCRUD/schemas"
	"github.com/gin-gonic/gin"
)

func ListProducts(ctx *gin.Context) {
	products := []schemas.Product{}

	if err := db.Find(&products).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing products")
		return
	}

	sendSuccess(ctx, "list-products", products)
}

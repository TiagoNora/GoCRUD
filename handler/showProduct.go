package handler

import (
	"net/http"

	"github.com/TiagoNora/GoCRUD/schemas"
	"github.com/gin-gonic/gin"
)

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

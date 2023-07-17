package handler

import (
	"net/http"

	"github.com/TiagoNora/GoCRUD/schemas"
	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {

	request := CreateProductRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	product := schemas.Product{
		Description: request.Description,
		Designation: request.Designation,
		Company:     request.Company,
		Price:       request.Price,
	}

	if err := db.Create(&product).Error; err != nil {
		logger.Errorf("error creating product: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-product", product)
}

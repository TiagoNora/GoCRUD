package handler

import (
	"net/http"

	"github.com/TiagoNora/GoCRUD/schemas"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

// @BasePath /api/v1

// @Summary Update product
// @Description Update a product
// @Tags Products
// @Accept json
// @Produce json
// @Param id query string true "Product Identification"
// @Param opening body UpdateProductRequest true "Opening data to Update"
// @Success 200 {object} UpdateProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /secured/product [put]
func UpdateProduct(ctx *gin.Context) {
	request := UpdateProductRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Product{}

	if err := db.First(&opening, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}
	// Update opening
	if request.Description != "" {
		opening.Description = request.Description
	}

	if request.Designation != "" {
		opening.Designation = request.Designation
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Price.GreaterThan(decimal.NewFromInt(0)) {
		opening.Price = request.Price
	}

	// Save opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}
	sendSuccess(ctx, "update-opening", opening)
}

package handler

import (
	"net/http"

	"github.com/TiagoNora/GoCRUD/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create product
// @Description Create a new product
// @Tags Products
// @Accept json
// @Produce json
// @Param request body CreateProductRequest true "Request body"
// @Success 200 {object} CreateProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /product [post]
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

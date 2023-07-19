package handler

import (
	"net/http"

	"github.com/TiagoNora/GoCRUD/auth"
	"github.com/TiagoNora/GoCRUD/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Generate a token
// @Description Generate a new token
// @Tags Users
// @Accept json
// @Produce json
// @Param request body TokenRequest true "Request body"
// @Success 200 {string} string
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/token [post]
func GenerateToken(ctx *gin.Context) {
	request := TokenRequest{}
	ctx.BindJSON(&request)

	user := schemas.User{}

	if err := db.First(&user, "email = ?", request.Email).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	if err := user.CheckPassword(request.Password); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	tokenString, err := auth.GenerateJWT(user.Email, user.Username, user.Role)
	if err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	sendSuccess(ctx, "create-jwt", tokenString)

}

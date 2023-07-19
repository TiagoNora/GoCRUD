package handler

import (
	"net/http"

	"github.com/TiagoNora/GoCRUD/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create a user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "Request body"
// @Success 200 {object} CreateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user [post]
func CreateUser(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	userSearched := schemas.User{}

	result := db.Where("username = ? OR email = ?", user.Username, user.Email).Find(&userSearched)

	if result.RowsAffected > 0 {
		sendError(ctx, http.StatusBadRequest, "already exists a user with that username or email")
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-user", user)
}

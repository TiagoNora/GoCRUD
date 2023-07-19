package handler

import (
	"net/http"

	"github.com/TiagoNora/GoCRUD/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update user
// @Description Update a user
// @Security bearerToken
// @Tags Users
// @Accept json
// @Produce json
// @Param id query string true "User Identification"
// @Param opening body UpdateUserRequest true "Opening data to Update"
// @Success 200 {object} UpdateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /secured/user [put]
func UpdateUser(ctx *gin.Context) {
	request := UpdateUserRequest{}

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

	user := schemas.User{}
	userSearched := schemas.User{}

	if err := db.First(&user, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}
	user.Name = request.Name
	user.Username = request.Username
	user.Email = request.Email
	user.Password = request.Password

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

	// Save opening
	if err := db.Save(&user).Error; err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}
	sendSuccess(ctx, "update-opening", user)
}

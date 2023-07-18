package handler

import (
	"fmt"
	"net/http"

	"github.com/TiagoNora/GoCRUD/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteUser(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	user := schemas.User{}
	// Find Opening
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("user with id: %s not found", id))
		return
	}
	// Delete Opening
	if err := db.Delete(&user).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting user with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-user", user)
}

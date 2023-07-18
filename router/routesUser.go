package router

import (
	"github.com/TiagoNora/GoCRUD/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutesUser(v1 *gin.RouterGroup) {
	v1.GET("/users", handler.ListUsers)
	v1.GET("/user", handler.ShowUser)
	v1.POST("/user", handler.CreateUser)
	v1.DELETE("/user", handler.DeleteUser)
	v1.PUT("/user", handler.UpdateUser)
	v1.POST("/user/token", handler.GenerateToken)
}

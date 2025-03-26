package user

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(router *gin.Engine, userHandler *UserHandler) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/:email", userHandler.GetUserByEmail)
	}
}

package auth

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(router *gin.Engine, authHandler *AuthHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}
}

package main

import (
	"screaming-architecture-skeleton/auth"
	"screaming-architecture-skeleton/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize authentication components
	authRepo := auth.NewAuthRepository()
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService)

	// Initialize user components
	userRepo := user.NewUserRepository()
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// Register routes
	auth.RegisterAuthRoutes(router, authHandler)
	user.RegisterUserRoutes(router, userHandler)

	router.Run(":8080")
}

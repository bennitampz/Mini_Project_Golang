package routes

import (
	"tokoku_go/handlers"
	"tokoku_go/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v2")

	// Authentication routes
	api.Post("/register", handlers.Register)
	api.Post("/login", handlers.Login)

	// Protected routes
	api.Use(middleware.AuthMiddleware())
	api.Get("/profile", handlers.GetProfile)
	api.Post("/logout", handlers.Logout)
}

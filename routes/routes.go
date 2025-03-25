package routes

import (
	"mini-project-BE-Evermos/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// api := app.Group("/api")
	auth := app.Group("/auth")

	auth.Post("/register", handlers.Register)

}
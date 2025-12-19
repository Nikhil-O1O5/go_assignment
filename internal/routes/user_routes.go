package routes

import (
	"go-backend-task/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	// Grouping routes under /users as per requirements [cite: 31, 44, 53, 66, 70]
	api := app.Group("/users")

	api.Post("/", userHandler.CreateUser)     // Create [cite: 31]
	api.Get("/", userHandler.ListUsers)      // List All [cite: 70]
	api.Get("/:id", userHandler.GetUser)     // Get by ID [cite: 44]
	api.Put("/:id", userHandler.UpdateUser)  // Update [cite: 53]
	api.Delete("/:id", userHandler.DeleteUser) // Delete [cite: 66]
}
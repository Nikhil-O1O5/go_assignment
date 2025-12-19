package main

import (
	"go-backend-task/config"
	"go-backend-task/internal/handler"
	"go-backend-task/internal/logger"    
	"go-backend-task/internal/repository"
	"go-backend-task/internal/routes"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	logger.InitLogger() 
	defer logger.Log.Sync()

	dbConn, err := config.InitDB()
	if err != nil {
		logger.Log.Fatal("Failed to connect to database")
	}
	defer dbConn.Close()

	userRepo := repository.NewUserRepository(dbConn)
	userHandler := handler.NewUserHandler(userRepo)
	
	app := fiber.New()

	app.Use(requestid.New())

	app.Use(fiberLogger.New(fiberLogger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path} ID=${header:X-Request-ID}\n",
	}))

	routes.SetupUserRoutes(app, userHandler)

	logger.Log.Info("Server starting on port 3000")
	if err := app.Listen(":3000"); err != nil {
		logger.Log.Fatal("Server failed to start")
	}
}
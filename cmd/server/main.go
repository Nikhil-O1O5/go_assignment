package main

import (
	"fmt"
	"go-backend-task/config"
	"go-backend-task/internal/logger"
	"os"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger()
	defer logger.Log.Sync()

	logger.Log.Info("Starting the application...")

	db, err := config.InitDB()
	if err != nil {
		logger.Log.Error("Database connection failed", zap.Error(err))
		os.Exit(1)
	}
	defer db.Close()

	logger.Log.Info("Successfully connected to the database!")
	fmt.Println("--- Connection Test Passed ---")
}
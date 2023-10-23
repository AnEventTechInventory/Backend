package main

import (
	"github.com/AnEventTechInventory/Backend/pkg/arguments"
	"github.com/AnEventTechInventory/Backend/pkg/configConstants"
	"github.com/AnEventTechInventory/Backend/pkg/database"
	"github.com/AnEventTechInventory/Backend/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func startServer() {
	// Parse command-line arguments
	arguments.ParseArgs()

	// Check if the data folder is present
	if _, err := os.Stat(configConstants.DataFolder); os.IsNotExist(err) {
		// Create the data folder if it doesn't exist
		err := os.Mkdir(configConstants.DataFolder, os.ModePerm)
		if err != nil {
			return
		}
	}

	logger.InitLogger()
	logger.Get().Println("Starting the application...")
	if !database.InitDatabase() {
		logger.Get().Println("Failed to initialize database")
		os.Exit(-1)
		return
	}

	// Create a channel to listen for signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	// Handle signals
	go func() {
		<-sigChan
		// Perform graceful shutdown or cleanup actions here
		stopServer()
		// For example, save data, close connections, etc.
		// Then, exit the application
		os.Exit(0)
	}()
}

func stopServer() {
	dbInstance, _ := database.Database.DB()
	err := dbInstance.Close()
	if err != nil {
		logger.Get().Printf("There was an error closing the database connection:%v\n", err)
		return
	}
	logger.Get().Println("Stopping the application...")
}

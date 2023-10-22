package main

import (
	"os"
	"os/signal"
	"syscall"
)

var DataFolder = "data"

func startServer() {
	// Parse command-line arguments
	parseArgs()

	// Check if the data folder is present
	if _, err := os.Stat(DataFolder); os.IsNotExist(err) {
		// Create the data folder if it doesn't exist
		err := os.Mkdir(DataFolder, os.ModePerm)
		if err != nil {
			return
		}
	}

	initLogger()
	initDatabase()

	Logger.Println("Starting the application...")

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
	dbInstance, _ := Database.DB()
	err := dbInstance.Close()
	if err != nil {
		Logger.Printf("There was an error closing the database connection:%v\n", err)
		return
	}
	Logger.Println("Stopping the application...")
}

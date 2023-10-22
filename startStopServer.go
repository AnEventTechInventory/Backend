package main

import (
	"os"
	"os/signal"
	"syscall"
)

func startServer() {
	// Parse command-line arguments
	parseArgs()

	loggerStart()

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

	// Check if the data folder is present
	DataFolder := "data"
	if _, err := os.Stat(DataFolder); os.IsNotExist(err) {
		// Create the data folder if it doesn't exist
		Logger.Println("Creating data folder...")
		os.Mkdir(DataFolder, os.ModePerm)
	}
}

func stopServer() {
	Logger.Println("Stopping the application...")
}

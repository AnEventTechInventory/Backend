package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Create a channel to listen for signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	// Handle signals
	go func() {
		<-sigChan
		// Perform graceful shutdown or cleanup actions here
		StopServer()
		// For example, save data, close connections, etc.
		// Then, exit the application
		os.Exit(0)
	}()

	// Your application code here
}

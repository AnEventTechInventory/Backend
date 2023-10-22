package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger

func initLogger() {
	logDir := filepath.Join(DataFolder, "log")

	// If the log directory does not exist, create it
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			log.Fatal(err)
		}
	}

	logFileName := filepath.Join(logDir, fmt.Sprintf("app-%s.log", time.Now().Format("2006-01-02-15-04-05")))
	logFile, err := os.Create(logFileName)
	if err != nil {
		log.Fatalf("Error opening log file: %v\n", err)
	}

	// Create a new Logrus logger instance
	Logger = log.New()
	Logger.Out = logFile // Set the output to the log file

	// If Args.verbose is true, set the log level to Debug
	if Args.verbose {
		Logger.SetLevel(log.DebugLevel)
	} else {
		Logger.SetLevel(log.DebugLevel)
	}
}

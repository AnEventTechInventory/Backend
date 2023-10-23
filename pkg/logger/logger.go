package logger

import (
	"fmt"
	"github.com/AnEventTechInventory/Backend/pkg/configConstants"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"time"
)

var logger *log.Logger = nil

func Get() *log.Logger {
	if logger == nil {
		InitLogger()
	}
	return logger
}

func InitLogger() {
	// Check if the logger was already initialized
	if logger != nil {
		logger.Println("Logger already initialized")
		return
	}

	logDir := filepath.Join(configConstants.DataFolder, "log")

	// If the log directory does not exist, create it
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			_ = fmt.Sprintf("Error creating log directory: %v\n", err)
		}
	}

	logFileName := filepath.Join(logDir, fmt.Sprintf("app-%s.log", time.Now().Format("2006-01-02-15-04-05")))
	logFile, err := os.Create(logFileName)
	if err != nil {
		_ = fmt.Sprintf("Error creating log file: %v\n", err)
	}

	// Create a new Logrus logger instance
	logger = log.New()
	logger.SetOutput(io.MultiWriter(os.Stdout, logFile))

	// If Args.verbose is true, set the log level to Debug
	/* if arguments.Args.Verbose {
		logger.SetLevel(log.DebugLevel)
	} else {
		logger.SetLevel(log.DebugLevel)
	}*/

	logger.Println("Logger initialized")
}

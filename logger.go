package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

var Logger *log.Logger

func loggerStart() {

	// if the log directory does not exist, create it
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		err := os.Mkdir("log", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	logFileName := filepath.Join("log", fmt.Sprintf("app-%s.log", time.Now().Format("2006-01-02-12H-04M-05S")))
	logFile, err := os.Create(logFileName)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		_, err := fmt.Fprintln(os.Stdout, "Error opening log file: %v", err)
		if err != nil {
			return
		}
		os.Exit(1) // Exit the program with a non-zero exit code
	}

	// Create a multi-writer to write to both the file and os.Stdout (terminal)
	multiWriter := io.MultiWriter(logFile, os.Stdout)

	// Create a logger with a prefix and flags

	if Args.verbose {
		Logger = log.New(multiWriter, "", log.Lshortfile|log.Ldate|log.Ltime)
	} else {
		Logger = log.New(multiWriter, "", log.Ldate|log.Ltime)
	}
}

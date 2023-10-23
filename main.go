package main

import (
	"github.com/AnEventTechInventory/Backend/pkg/httpServer"
)

func main() {
	startServer()

	// Your application code here
	httpServer.RunHttpServer()

	stopServer()
}

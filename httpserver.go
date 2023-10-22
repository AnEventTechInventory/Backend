package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func runHttpServer() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "5678"
		log.Printf("Defaulting to port %s", port)
	}

	r := gin.Default()

	// Add localhost to trusted proxies using the engine
	err := r.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	if err != nil {
		Logger.Fatal(err)
		return
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	runErr := r.Run(":" + port)
	if runErr != nil {
		Logger.Fatal(runErr)
		return
	}
}

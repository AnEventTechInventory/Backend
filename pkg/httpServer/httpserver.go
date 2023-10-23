package httpServer

import (
	"fmt"
	"github.com/AnEventTechInventory/Backend/api"
	"github.com/AnEventTechInventory/Backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"os"
)

func RunHttpServer() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "5678"
		logger.Logger.Printf("Defaulting to port %s", port)
	}

	r := gin.Default()

	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: logger.Logger.Writer(),
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("[%s] %s %s %s %d %s\n",
				params.TimeStamp.Format("2006/01/02 - 15:04:05"),
				params.Method,
				params.Path,
				params.Request.Proto,
				params.StatusCode,
				params.Latency,
			)
		},
	}))

	// Add loop-back IPs to trusted proxies using the engine
	err := r.SetTrustedProxies([]string{"::1", "127.0.0.1"})
	if err != nil {
		logger.Logger.Fatal(err)
		return
	}

	api.AddAllRoutes(r)

	runErr := r.Run(":" + port)
	if runErr != nil {
		logger.Logger.Fatal(runErr)
		return
	}
}

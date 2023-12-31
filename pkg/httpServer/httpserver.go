package httpServer

import (
	"fmt"
	"github.com/AnEventTechInventory/Backend/pkg/api"
	"github.com/AnEventTechInventory/Backend/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"os"
)

func RunHttpServer() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "5678"
		logger.Get().Printf("Defaulting to port %s", port)
	}

	r := gin.Default()

	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: logger.Get().Writer(),
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
		logger.Get().Fatal(err)
		return
	}

	// Add cors support
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
		ExposeHeaders:    []string{"Content-Length"},
	}))

	api.AddAllRoutes(r)

	store := cookie.NewStore([]byte(os.Getenv("COOKIE_SECRET")))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(csrf.Middleware(csrf.Options{
		Secret: os.Getenv("CSRF_SECRET"),
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	r.GET("/protected", func(c *gin.Context) {
		c.String(200, csrf.GetToken(c))
	})

	r.POST("/protected", func(c *gin.Context) {
		c.String(200, "CSRF token is valid")
	})

	runErr := r.Run(":" + port)
	if runErr != nil {
		logger.Get().Fatal(runErr)
		return
	}
}

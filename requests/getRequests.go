package requests

import "github.com/gin-gonic/gin"

func (c *gin.Context) HomePage() {

}

func (c *gin.Context) HealthCheck() {

}

func AddAllRoutes(r *gin.Engine) {
	r.GET("/", HomePage)

	r.GET("/api/v1/health", HealthCheck)
}

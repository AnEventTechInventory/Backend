package api

import (
	"github.com/AnEventTechInventory/Backend/pkg/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func homePage(context *gin.Context) {
	context.String(http.StatusOK, "Hello World!")
}

func apiHandler(context *gin.RouterGroup) {
	context.GET("/v:version", func(c *gin.Context) {
		versionString := c.Param("version")
		version, err := strconv.Atoi(versionString)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid version"})
			return
		}

		switch version {
		case 1:
			versionOne := context.Group("/v1") // Use the 'context' here to create the group.
			v1.Handler(versionOne)
		}
	})
}

func AddAllRoutes(context *gin.Engine) {

	context.GET("", homePage)

	apiGroup := context.Group("/api")
	apiHandler(apiGroup)
}

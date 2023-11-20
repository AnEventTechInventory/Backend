package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func homePage(context *gin.Context) {
	context.String(http.StatusOK, "Hello World!")
}

func AddAllRoutes(context *gin.Engine) {

	context.GET("", homePage)
	RegisterDevices(context)
	RegisterManufacturers(context)
	RegisterLocations(context)
}

package api

import (
	"github.com/gin-gonic/gin"
)

type requestInterface interface {
	list(context *gin.Context)
	get(context *gin.Context)
	create(context *gin.Context)
	update(context *gin.Context)
	delete(context *gin.Context)
}

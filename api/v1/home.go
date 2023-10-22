package v1

import "github.com/gin-gonic/gin"

func Handler(context *gin.RouterGroup) {

	// For now just reflect the path back to the user.
	context.Any("/*path", func(c *gin.Context) {
		c.JSON(200, gin.H{"path": c.Param("path")})
	})
}

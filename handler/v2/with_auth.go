package v2

import (
	"github.com/gin-gonic/gin"
	"ynufesCounterBackend/domain"
	"ynufesCounterBackend/middleware"
)

func WithAuth(process func(userID domain.UserID, c *gin.Context), c *gin.Context) {
	userID := c.GetInt(middleware.ContextAuthID)
	if userID == 0 {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}
	process(domain.UserID(userID), c)
}

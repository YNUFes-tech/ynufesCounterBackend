package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"ynufesCounterBackend/pkg/jwt"
	"ynufesCounterBackend/pkg/setting"
)

const ContextAuthID = "auth_id"

type AuthMiddleware struct {
	jwtSecret string
}

func NewAuthMiddleware() *AuthMiddleware {
	config := setting.Get()
	return &AuthMiddleware{
		jwtSecret: config.Service.Authentication.JWTSecret,
	}
}

func (m AuthMiddleware) Handle(c *gin.Context) {
	header := c.GetHeader("Authorization")
	// validate Bearer token
	if len(header) < 8 || header[:7] != "Bearer " {
		c.AbortWithStatus(401)
		return
	}

	token := header[7:]
	verification, err := jwt.Verify(token, m.jwtSecret)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"message": "invalid token"})
		return
	}
	// check if Id is convertable as int
	id, err := strconv.Atoi(verification.Id)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"message": "invalid token"})
		return
	}
	c.Set(ContextAuthID, id)
	c.Next()
}

package v2

import (
	"github.com/gin-gonic/gin"
	"ynufesCounterBackend/pkg/line"
	"ynufesCounterBackend/pkg/setting"
)

type LineAuthHandler struct {
	verifier line.AuthVerifier
}

func NewLineAuthHandler() LineAuthHandler {
	config := setting.Get()
	return LineAuthHandler{
		verifier: line.NewAuthVerifier(
			config.ThirdParty.Line.CallbackURI, config.ThirdParty.Line.ClientID, config.ThirdParty.Line.ClientSecret),
	}
}

func (h LineAuthHandler) Handle(c *gin.Context) {
	code := c.Request.URL.Query().Get("code")
	state := c.Request.URL.Query().Get("state")
	if code == "" || state == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "bad request, code and state are required",
		})
		return
	}
	profile, err := h.verifier.RequestAccessToken(code, state)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "ok",
		"profile": profile,
	})
}

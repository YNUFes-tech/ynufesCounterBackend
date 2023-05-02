package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"ynufesCounterBackend/handler"
	"ynufesCounterBackend/pkg/firebase"
)

func main() {
	engine := gin.New()

	// endpoints that requires authentication
	apiV1 := engine.Group("/api/v1")
	apiV1.Handle("GET", "/hello", handleHello)

	//lineAuthHandler := handler.NewLineAuthHandler()
	//apiV1.Handle("GET", "/line/auth/callback", lineAuthHandler.Handle)

	//authMiddleware := middleware.NewAuthMiddleware()
	//authRg := apiV1.Use(authMiddleware.Handle)

	fb := firebase.New()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowMethods("OPTIONS")
	apiV1.Use(cors.New(corsConfig))
	countHandler := handler.NewCountHandler(fb)
	apiV1.Handle("POST", "/count/entry", countHandler.HandleEntry)
	apiV1.Handle("POST", "/count/exit", countHandler.HandleExit)

	if err := engine.Run(":8080"); err != nil {
		return
	}
}

func handleHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, world!",
	})
}

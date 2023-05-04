package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	hV1 "ynufesCounterBackend/handler/v1"
	hV2 "ynufesCounterBackend/handler/v2"
	"ynufesCounterBackend/middleware"
	"ynufesCounterBackend/pkg/firebase"
)

func main() {
	engine := gin.New()

	// endpoints that requires authentication
	apiV1 := engine.Group("/api/v1")
	implementAPIV1(apiV1)

	apiV2 := engine.Group("/api/v2")
	implementAPIV2(apiV2)
	//lineAuthHandler := handler.NewLineAuthHandler()
	//apiV1.Handle("GET", "/line/auth/callback", lineAuthHandler.Handle)

	// authMiddleware := middleware.NewAuthMiddleware()
	//authRg := apiV1.Use(authMiddleware.Handle)

	if err := engine.Run(":8080"); err != nil {
		return
	}
}

func handleHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, world!",
	})
}

func implementAPIV1(rg *gin.RouterGroup) {
	rg.Handle("GET", "/hello", handleHello)
	fb := firebase.New()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowMethods("OPTIONS")
	rg.Use(cors.New(corsConfig))
	countHandler := hV1.NewCountHandler(fb)
	rg.Handle("POST", "/count/entry", countHandler.HandleEntry)
	rg.Handle("POST", "/count/exit", countHandler.HandleExit)
}

func implementAPIV2(rg *gin.RouterGroup) {
	authMiddleware := middleware.NewAuthMiddleware()
	rg.Use(authMiddleware.Handle)
	fb := firebase.New()
	countHandler := hV2.NewCountHandler(fb)
	rg.Handle("POST", "/count/entry", countHandler.HandleEntry)
}

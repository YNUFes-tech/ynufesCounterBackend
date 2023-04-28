package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()

	engine.Handle("GET", "/hello", handleHello)

	if err := engine.Run(":8080"); err != nil {
		return
	}
}

func handleHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, world!",
	})
}

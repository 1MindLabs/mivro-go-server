package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func rootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "We're cooking up something great! Check back soon.",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", rootHandler)

	if err := router.Run(":8080"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

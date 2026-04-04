package main

import (
	"github.com/gin-gonic/gin"

	"unit-api/internal/database"
)

func main() {
	database.Connect()
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server is running",
		})
	})

	router.Run(":8080")
}
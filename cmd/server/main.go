package main

import (
	"github.com/gin-gonic/gin"

	"unit-api/internal/database"
	"unit-api/internal/handlers"
)

func main() {
	database.Connect()
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/units", handlers.GetUnitsHandler)
		api.GET("/units/:id", handlers.GetUnitByIDHandler)
		api.POST("/units", handlers.CreateUnitHandler)
		api.PUT("/units/:id", handlers.UpdateUnitStatusHandler)
	}

	router.Run(":8080")
}
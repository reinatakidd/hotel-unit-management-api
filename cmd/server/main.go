package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"unit-api/internal/database"
	"unit-api/internal/handlers"
)

func main() {
    database.Connect()
    router := gin.Default()

    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    api := router.Group("/api")
    {
        api.GET("/units", handlers.GetUnitsHandler)
        api.GET("/units/:id", handlers.GetUnitByIDHandler)
        api.POST("/units", handlers.CreateUnitHandler)
        api.PUT("/units/:id", handlers.UpdateUnitStatusHandler)
    }

    router.Run(":8080")
}
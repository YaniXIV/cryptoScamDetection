package server

import (
	"cryptoScamDetection/config"
	"cryptoScamDetection/handlers"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.Default()
	r.Use(cors.Default())
	fmt.Println("Starting the server!")
	r.POST("/analyze", handlers.HandleAnalyze)

	port := config.GetEnv("PORT", "8080")
	r.Run(":" + port)
}

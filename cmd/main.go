package main

import (
	"log"
	"os"

	"ticket-system/config"
	"ticket-system/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	config.ConnectDatabase()

	router := gin.Default()

	routes.SetupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
package main

import (
	"log"
	"os"

	"github.com/agrimsharma2/golang-web-service/handlers"
	"github.com/agrimsharma2/golang-web-service/models"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	// Initialize routes
	models.Init() // Initialize models

	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:id", handlers.GetBookByID)
	router.POST("/books", handlers.CreateBook)
	router.PUT("/books/:id", handlers.UpdateBook)
	router.DELETE("/books/:id", handlers.DeleteBook)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

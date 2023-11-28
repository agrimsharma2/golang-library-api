package handlers

import (
	"net/http"

	"github.com/agrimsharma2/golang-web-service/models"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.GetBooks())
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := models.GetBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// TODO: Implement CreateBook, UpdateBook, and DeleteBook handlers

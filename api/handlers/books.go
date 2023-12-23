package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricsjs/infohub-api/api/models"
	"github.com/ricsjs/infohub-api/api/repository"
)

func GetAllBooks(c *gin.Context) []models.Book {
	books, err := repository.GetAllBooks()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting books from database"})
		return nil
	}

	return books
}

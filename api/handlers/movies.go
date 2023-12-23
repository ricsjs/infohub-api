package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricsjs/infohub-api/api/models"
	"github.com/ricsjs/infohub-api/api/repository"
)

func GetAllMovies(c *gin.Context) []models.Movie {
	movies, err := repository.GetAllMovies()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting movies from database"})
		return nil
	}

	return movies
}

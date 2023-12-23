package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricsjs/infohub-api/api/models"
	"github.com/ricsjs/infohub-api/api/repository"
)

func GetAllBeers(c *gin.Context) []models.Beer { // Assuming `GetAllBeers` returns a slice of `Beer`
	beers, err := repository.GetAllBeers()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting beers from database"})
		return nil // Return an empty slice on error
	}

	return beers // Return the retrieved beers
}

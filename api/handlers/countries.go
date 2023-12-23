package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricsjs/infohub-api/api/models"
	"github.com/ricsjs/infohub-api/api/repository"
)

func GetAllCountries(c *gin.Context) []models.Country {
	countries, err := repository.GetAllCountries()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting countries from database"})
		return nil
	}

	return countries
}

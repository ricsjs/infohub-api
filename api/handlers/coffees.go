package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricsjs/infohub-api/api/models"
	"github.com/ricsjs/infohub-api/api/repository"
)

func GetAllCoffee(c *gin.Context) []models.Coffee {
	coffees, err := repository.GetAllCoffees()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting coffees from database"})
		return nil
	}

	return coffees
}

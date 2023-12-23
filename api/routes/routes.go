package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricsjs/infohub-api/api/handlers"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to InfoHub API! InfoHub is an API developed in Golang designed to be a central hub of information covering a variety of topics. Furthermore, InfoHub offers a reliable and affordable source for students and enthusiasts who want to explore and consume diverse data.")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/beer", func(c *gin.Context) {
			c.JSON(http.StatusOK, handlers.GetAllBeers(c))
		})

		v1.GET("/book", func(c *gin.Context) {
			c.JSON(http.StatusOK, handlers.GetAllBooks(c))
		})

		v1.GET("/coffee", func(c *gin.Context) {
			c.JSON(http.StatusOK, handlers.GetAllCoffee(c))
		})

		v1.GET("/country", func(c *gin.Context) {
			c.JSON(http.StatusOK, handlers.GetAllCountries(c))
		})

		v1.GET("/movie", func(c *gin.Context) {
			c.JSON(http.StatusOK, handlers.GetAllMovies(c))
		})
	}
}

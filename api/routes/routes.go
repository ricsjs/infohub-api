package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricsjs/infohub-api/api/handlers"
)

func SetupRoutes(r *gin.Engine) {

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

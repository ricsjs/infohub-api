package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		v1.GET("/beer", func(c *gin.Context) {
			c.JSON(http.StatusOK, "beer")
		})

		v1.GET("/book", func(c *gin.Context) {
			c.JSON(http.StatusOK, "book")
		})

		v1.GET("/coffee", func(c *gin.Context) {
			c.JSON(http.StatusOK, "coffee")
		})

		v1.GET("/country", func(c *gin.Context) {
			c.JSON(http.StatusOK, "country")
		})

		v1.GET("/movie", func(c *gin.Context) {
			c.JSON(http.StatusOK, "movie")
		})
	}
}

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shlason/url-shortener/controllers"
)

func RegisteShortRoutes(r *gin.Engine) {
	r.POST("/short", controllers.CreateShortURL)
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte("Hi"))
	})
}

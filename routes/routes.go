package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisteStaticContentRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte("Hi Index"))
	})
}

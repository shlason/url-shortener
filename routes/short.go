package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shlason/url-shortener/controllers"
)

func RegisteShortRoutes(r *gin.RouterGroup) {
	r.POST("/short", controllers.CreateShortURL)
	r.GET("/:shortID", controllers.GetShortIDRediect)
}

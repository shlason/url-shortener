package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shlason/url-shortener/configs"
)

func RegisteStaticContentRoutes(r *gin.Engine) {
	r.Static("/static", fmt.Sprintf("%s/static", configs.Server.FeWrokDir))
	r.StaticFile("/", fmt.Sprintf("%s/index.html", configs.Server.FeWrokDir))
	r.StaticFile("/notfound", fmt.Sprintf("%s/index.html", configs.Server.FeWrokDir))
}

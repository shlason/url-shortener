package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shlason/url-shortener/configs"
	"github.com/shlason/url-shortener/docs"
	"github.com/shlason/url-shortener/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

// @title           URL-Shortener Example API
// @version         1.0
// @description     This is a sample server celler server.

// @contact.name   sidesideeffect.io
// @contact.url    https://github.com/shlason/url-shortener
// @contact.email  nocvi111@gmail.com

// @license.name  MIT
// @license.url   https://github.com/shlason/url-shortener/blob/main/LICENSE

// @host      short.sidesideeffect.io
// @BasePath  /api
func main() {
	var g errgroup.Group

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://*.short.sidesideeffect.io"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin"},
	}))

	apiRoute := r.Group("/api")

	routes.RegisteStaticContentRoutes(r)
	routes.RegisteShortRoutes(apiRoute)

	docs.SwaggerInfo.Schemes = []string{"https"}

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// r.Run()

	g.Go(func() error {
		return http.ListenAndServe(":http", http.RedirectHandler(fmt.Sprintf("https://%s", configs.Server.Host), http.StatusSeeOther))
	})
	g.Go(func() error {
		return http.Serve(autocert.NewListener(configs.Server.Host), r)
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

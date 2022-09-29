package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shlason/url-shortener/configs"
	"github.com/shlason/url-shortener/routes"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group

	r := gin.Default()

	apiRoute := r.Group("/api")

	routes.RegisteStaticContentRoutes(r)
	routes.RegisteShortRoutes(apiRoute)

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

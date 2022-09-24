package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shlason/url-shortener/routes"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group
	r := gin.Default()

	routes.RegisteShortRoutes(r)

	g.Go(func() error {
		return http.ListenAndServe(":http", http.RedirectHandler("https://sidesideeffect.io", http.StatusSeeOther))
	})
	g.Go(func() error {
		return http.Serve(autocert.NewListener("sidesideeffect.io"), r)
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

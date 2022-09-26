package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shlason/url-shortener/routes"
)

func main() {
	// var g errgroup.Group
	r := gin.Default()

	apiRoute := r.Group("/api")

	routes.RegisteStaticContentRoutes(r)
	routes.RegisteShortRoutes(apiRoute)

	r.Run()

	// g.Go(func() error {
	// 	return http.ListenAndServe(":http", http.RedirectHandler("https://sidesideeffect.io", http.StatusSeeOther))
	// })
	// g.Go(func() error {
	// 	return http.Serve(autocert.NewListener("sidesideeffect.io"), r)
	// })

	// if err := g.Wait(); err != nil {
	// 	log.Fatal(err)
	// }
}

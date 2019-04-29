package server

import (
	"net/http"
)

// Run run server
func Run(addr string) error {
	// build router
	routes := buildRoutes()

	server := http.Server{
		Addr:    addr,
		Handler: routes,
	}

	return server.ListenAndServe()
}

package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/zuhrulumam/learning-go/pkg/users"
)

func buildRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {

		r.Route("/users", func(r chi.Router) {
			users.Routes(r)
		})

	})

	log.Println("List Routes : ")
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	return r
}

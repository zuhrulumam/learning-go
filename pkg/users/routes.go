package users

import (
	// "log"
	// "net/http"

	"github.com/go-chi/chi"
	// "github.com/zuhrulumam/learning-go/pkg/database"
)

// Routes list user routes
func Routes(r chi.Router) {
	r.Get("/", v1ListUsersHandler)
	r.Post("/", v1CreateUserHandler)
	r.Put("/{id}", v1UpdateUsersHandler)
	r.Delete("/{id}", v1DeleteUsersHandler)
}

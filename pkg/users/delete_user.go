package users

import (
	"log"
	"net/http"

	"github.com/zuhrulumam/learning-go/pkg/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func v1DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("deleting user")

	id := chi.URLParam(r, "id")

	query := `
		delete from users where id = ?
	`
	err := database.Exec(query, id)
	if err != nil {
		log.Println("error creating " + err.Error())
	}

	render.Status(r, 200)
	render.JSON(w, r, "success deleting user")
}

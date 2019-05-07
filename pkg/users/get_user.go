package users

import (
	"net/http"

	"github.com/zuhrulumam/learning-go/pkg/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func v1GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	query := `
		select name, address from users where id = ?
	`

	var user User
	err := database.Query(&user, query, id)
	if err != nil {
		render.Status(r, 500)
		render.JSON(w, r, "error on query "+err.Error())
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, &user)
}

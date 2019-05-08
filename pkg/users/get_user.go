package users

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/zuhrulumam/learning-go/pkg/models"
)

func v1GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	intID, _ := strconv.Atoi(id)

	var u models.Users
	u.ID = uint(intID)
	err := u.GetUser()
	if err != nil {
		render.Status(r, 500)
		render.JSON(w, r, "error on query "+err.Error())
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, u)
}

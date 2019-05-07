package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zuhrulumam/learning-go/pkg/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func v1UpdateUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("updating user")

	id := chi.URLParam(r, "id")

	var req v1CreateUsersPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("error on reading json")
		return
	}

	query := `
		update users set 
		name = ?, 
		address = ?, 
		updated_at = now() 
		where id = ? 
	`

	err := database.Exec(query, req.Name, req.Address, id)
	if err != nil {
		log.Println("error creating " + err.Error())
	}

	render.Status(r, 200)
	render.JSON(w, r, "success updating user")
}

package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/zuhrulumam/learning-go/pkg/database"
)

type v1CreateUsersPayload struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func v1CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("creating user")

	var req v1CreateUsersPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("error on reading json")
		return
	}

	query := `
		insert into users (name, address, created_at) 
		values
		(?, ?, now())
	`
	err := database.Exec(query, req.Name, req.Address)
	if err != nil {
		log.Println("error creating " + err.Error())
	}

	render.Status(r, 201)
	render.JSON(w, r, "success creating user")
}

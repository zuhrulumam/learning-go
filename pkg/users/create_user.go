package users

import (
	"github.com/zuhrulumam/learning-go/pkg/database"
	"log"
	"net/http"
)

func v1CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("create user")
	query := `
		insert into users (name, address) 
		values 
		(?,?)
		returning * 
	`
	_, err := database.Exec(query, "umam", "test address")
	if err != nil {
		log.Println("error creating")
	}
}

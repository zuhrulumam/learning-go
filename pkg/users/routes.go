package users

import (
	// "log"
	// "net/http"

	"github.com/go-chi/chi"
	// "github.com/zuhrulumam/learning-go/pkg/database"
)

// Routes list user routes
func Routes(r chi.Router) {

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {

	// 	rows, err := database.Query(`select * from test`)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	log.Println(rows)

	// 	var (
	// 		id   uint
	// 		name string
	// 	)

	// 	type user struct {
	// 		id   uint
	// 		name string
	// 	}

	// 	var users []user
	// 	for rows.Next() {
	// 		err := rows.Scan(&id, &name)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		log.Println(id, name)
	// 		currUser := user{
	// 			id:   id,
	// 			name: name,
	// 		}
	// 		users = append(users, currUser)
	// 	}

	// 	log.Println(users)
	// 	log.Println("Get / from users")
	// })
	r.Get("/", v1ListUsersHandler)
	r.Post("/", v1CreateUserHandler)
	r.Put("/{id}", v1UpdateUsersHandler)
}

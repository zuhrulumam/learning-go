package users

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/zuhrulumam/learning-go/pkg/database"
)

func Routes(r chi.Router) {

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := database.Query(`select * from test`)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(rows)
		log.Println("Get / from users")
	})
}

package users

import (
	"net/http"

	"github.com/zuhrulumam/learning-go/pkg/database"

	"log"

	"github.com/go-chi/render"
)

func v1ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	// query := `
	// 	select name, (select count(id) from users) as countData from users
	// `

	// rows, err := database.Query(query)
	// if err != nil {
	// 	render.Status(r, 500)
	// 	render.JSON(w, r, "error on query "+err.Error())
	// 	return
	// }
	// defer rows.Close()

	// type User struct {
	// 	Name    string `json:"name"`
	// 	Address string `json:"address"`
	// }

	// var countData int
	// users := []User{}
	// for rows.Next() {
	// 	user := User{}
	// 	err = rows.Scan(&user.Name, &countData)
	// 	if err != nil {
	// 		render.Status(r, 500)
	// 		render.JSON(w, r, "error scanning")
	// 		return
	// 	}
	// 	users = append(users, user)
	// }

	// log.Println(countData)
	log.Println(w)
	query := `
		select name, address from users
	`

	countQuery := `
		select count(id) as countData from users
	`

	countRows, err := database.Query(countQuery)
	if err != nil {
		render.Status(r, 500)
		render.JSON(w, r, "error on query "+err.Error())
		return
	}
	defer countRows.Close()

	var countData uint
	for countRows.Next() {
		err = countRows.Scan(&countData)
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, "error scanning count"+err.Error())
			return
		}
	}

	rows, err := database.Query(query)
	if err != nil {
		render.Status(r, 500)
		render.JSON(w, r, "error on query "+err.Error())
		return
	}
	defer rows.Close()

	type User struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	users := []User{}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Name, &user.Address)
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, "error scanning")
			return
		}
		users = append(users, user)
	}

	log.Println(countData)

	render.Status(r, 200)
	render.JSON(w, r, &users)
}

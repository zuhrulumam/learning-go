package users

import (
	"log"
	"net/http"

	"github.com/zuhrulumam/learning-go/pkg/database"

	"github.com/go-chi/render"
)

// User model
type User struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// CountDataModel models
type CountDataModel struct {
	CountData uint `json:"count_data"`
}

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
	query := `
		select name, address from users
	`

	// column must be lowercase _ separated (like column in table name)
	countQuery := `
		select count(id) as count_data from users 
	`

	var countDataModel CountDataModel
	err := database.Query(&countDataModel, countQuery)
	if err != nil {
		render.Status(r, 500)
		render.JSON(w, r, "error on query "+err.Error())
		return
	}

	var users []User
	err = database.Query(&users, query)
	if err != nil {
		render.Status(r, 500)
		render.JSON(w, r, "error on query "+err.Error())
		return
	}

	log.Println("count data ", countDataModel.CountData)

	render.Status(r, 200)
	render.JSON(w, r, &users)
}

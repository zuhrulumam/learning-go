package main

import (
	"log"

	"github.com/zuhrulumam/learning-go/pkg/server"

	"github.com/zuhrulumam/learning-go/pkg/config"
	"github.com/zuhrulumam/learning-go/pkg/database"
)

func main() {
	// set config
	if err := config.ParseEnv(); err != nil {
		log.Println("Error Parsing Env")
	}

	// set database
	// if err := database.NewConnection(config.DbURL()); err != nil {
	// 	log.Println(err)
	// 	log.Println("Error Connection")
	// }

	database.Configure(config.DbURL())

	// rows, err := database.Query(`select * from test`)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(rows)

	// var (
	// 	id   uint
	// 	name string
	// )

	// type user struct {
	// 	id   uint
	// 	name string
	// }

	// var users []user
	// for rows.Next() {
	// 	err := rows.Scan(&id, &name)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(id, name)
	// 	currUser := user{
	// 		id:   id,
	// 		name: name,
	// 	}
	// 	users = append(users, currUser)
	// }

	// log.Println(users)

	// run server
	log.Println("Server run on ", config.ListenAddr())
	if err := server.Run(config.ListenAddr()); err != nil {
		log.Fatalln("Server Died")
	}
}

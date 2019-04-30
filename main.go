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

	database.Configure(config.DbURL())
	defer database.Quit()

	// run server
	log.Println("Server run on ", config.ListenAddr())
	if err := server.Run(config.ListenAddr()); err != nil {
		log.Fatalln("Server Died")
	}
}

package database

import (
	// "database/sql"
	"log"
)

func Configure(dbURL string) {

	conn, err := newConnection(dbURL)
	if err != nil {
		log.Println("Cannot connect database")
	}

	cs = conn

}

package database

import (
	"log"
)

// Quit database
func Quit() {
	log.Println("database closed")
	cs.quitFn()
}

// Configure configure connection
func Configure(dbURL string) {

	conn, err := newConnection(dbURL)
	if err != nil {
		log.Println("Cannot connect database")
	}

	cs = conn
}

// ConfigureTest configure connection test
func ConfigureTest(dbURL string) error {

	conn, err := newConnection(dbURL)
	if err != nil {
		log.Println("Cannot connect database")
		return err
	}

	cs = conn

	return err

}

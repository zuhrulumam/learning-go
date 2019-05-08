package models

import (
	"log"

	"github.com/zuhrulumam/learning-go/pkg/database"
)

// Users model
type Users struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// GetUser get user by id
func (u *Users) GetUser() (err error) {
	query := `
		select name, address from users where id = ?
	`
	err = database.Query(&u, query, u.ID)

	return
}

// CreateUser create user
func (u *Users) CreateUser() (err error) {
	query := `
		insert into users (name, address, created_at) 
		values
		(?, ?, now())
	`
	err = database.Exec(query, u.Name, u.Address)
	if err != nil {
		log.Println("error creating " + err.Error())
	}

	return
}

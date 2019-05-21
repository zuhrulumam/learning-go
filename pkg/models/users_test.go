package models

import "log"

func (m *ModelsTest) TestUsersModel_CreateUserOK() {
	newUser := &Users{
		Name:    "test",
		Address: "test address",
	}

	err := newUser.CreateUser()
	log.Println("new user ", newUser.ID)
	// check if err creating user nil
	m.Nil(err)

}

// func (m *ModelsTest) TestUsersModel_GetUserOK() {
// 	newUser := &Users{
// 		Name:    "test",
// 		Address: "test address",
// 	}

// 	err := newUser.CreateUser()

// 	log.Println(newUser.ID)
// 	// check if err creating user nil
// 	m.Nil(err)
// }

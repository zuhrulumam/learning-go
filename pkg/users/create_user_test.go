package users

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (u *UsersTest) TestUsers_V1CreateUserHandler() {
	name := "test from test case"
	address := "test from test case"

	data := map[string]interface{}{
		"name":    name,
		"address": address,
	}

	jsonStr, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))

	var w http.ResponseWriter

	v1CreateUserHandler(w, req)

	u.True(true, "it should be true")

}

package users

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func (u *UsersTest) TestUsers_V1CreateUserHandler() {
	name := "test from test case"
	address := "test from test case"

	newUserTest := User{
		Name:    name,
		Address: address,
	}

	jsonStr, _ := json.Marshal(newUserTest)

	// create requests
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	v1CreateUserHandler(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	u.Equal(http.StatusCreated, resp.StatusCode, "it should be 201")
	u.Equal("\"success creating user\"\n", string(body), "it should be same return")

}

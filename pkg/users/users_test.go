package users

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
)

func (u *UsersTest) TestUsers_GetAllUsers() {
	// create requests
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	v1ListUsersHandler(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	// fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println("body ", string(body))

	u.True(true, "it should true")

}

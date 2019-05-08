package users

import (
	"net/http"
	"net/http/httptest"
)

var newUserTest User

func (u *UsersTest) TestUsers_GetAllUsers() {
	// create requests
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	v1ListUsersHandler(w, req)
	resp := w.Result()
	// body, _ := ioutil.ReadAll(resp.Body)

	// fmt.Println(resp.Header.Get("Content-Type"))
	// fmt.Println("body ", string(body))

	u.Equal(http.StatusOK, resp.StatusCode, "it should be status ok")

}

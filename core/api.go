package core

import (
	"encoding/json"
	"net/http"
	"crypto/md5"
)

var users = make(map[string]string)
type User struct {
	Username string
	Password string
}

func OnPing(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"ok":   true,
		"pong": true,
	}

	responseBytes, _ := json.Marshal(response)
	_, _ = w.Write(responseBytes)
	return
}

func OnSignup(w http.ResponseWriter, r *http.Request) {
	var user *User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(user.Username) == 0 || len(user.Password) == 0  {
		http.Error(w, "Invalid params", http.StatusBadRequest)
		return
	}

	_, ok := users[user.Username]
	if ok {
		http.Error(w, "User " + user.Username + " already exists", http.StatusBadRequest)
		return
	}

	var hashedPass string
	var b [16]byte
	b = md5.Sum([]byte(user.Password))
	hashedPass = string(b[:])
	users[user.Username] = hashedPass
	_,_ = w.Write([]byte("User was added"))
}

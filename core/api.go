package core

import (
	"crypto/md5"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

var users = make(map[uuid.UUID]User)

type User struct {
	Username string `json:"username"`
	Password string
	UID      uuid.UUID `json:"uuid"`
}

type UserToReturn struct {
	Username string
	Uuid     uuid.UUID
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

	if len(user.Username) == 0 || len(user.Password) == 0 {
		http.Error(w, "Invalid params", http.StatusBadRequest)
		return
	}

	var hashedPass string
	var b [16]byte
	b = md5.Sum([]byte(user.Password))
	hashedPass = string(b[:])
	var newUser User
	newUser.Username = user.Username
	newUser.Password = hashedPass
	newUser.UID = uuid.New()
	users[newUser.UID] = newUser
	return
}

func OnGetUsers(w http.ResponseWriter, r *http.Request) {
	values := make([]UserToReturn, 0, len(users))
	for _, v := range users {
		userToReturn := CreateAUserToReturn(v)
		values = append(values, userToReturn)
	}

	responseBytes, _ := json.Marshal(values)
	_, _ = w.Write(responseBytes)
	return
}

func OnDeleteUser(w http.ResponseWriter, r *http.Request) {
	ss := strings.Split(r.URL.Path, "/")
	userUuid := ss[len(ss)-1]

	if len(userUuid) == 0 {
		http.Error(w, "Invalid query param", http.StatusBadRequest)
		return
	}

	parsedUuid, parseErr := uuid.Parse(userUuid)
	_, ok := users[parsedUuid]
	if parseErr != nil || !ok {
		http.Error(w, "User with uuid "+userUuid+" does not exist", http.StatusBadRequest)
		return
	}

	delete(users, parsedUuid)
	return
}

func OnGetUser(w http.ResponseWriter, r *http.Request) {
	ss := strings.Split(r.URL.Path, "/")
	userUuid := ss[len(ss)-1]

	if len(userUuid) == 0 {
		http.Error(w, "Invalid query param", http.StatusBadRequest)
		return
	}

	parsedUuid, parseErr := uuid.Parse(userUuid)
	user, ok := users[parsedUuid]
	if parseErr != nil || !ok {
		http.Error(w, "User with uuid "+userUuid+" does not exist", http.StatusBadRequest)
		return
	}

	responseBytes, _ := json.Marshal(CreateAUserToReturn(user))
	_, _ = w.Write(responseBytes)
}

func OnUpdateUser(w http.ResponseWriter, r *http.Request) {
	ss := strings.Split(r.URL.Path, "/")
	userUuid := ss[len(ss)-1]

	if len(userUuid) == 0 {
		http.Error(w, "Invalid query param", http.StatusBadRequest)
		return
	}

	parsedUuid, parseErr := uuid.Parse(userUuid)
	existingUser, ok := users[parsedUuid]
	if parseErr != nil || !ok {
		http.Error(w, "User with uuid "+userUuid+" does not exist", http.StatusBadRequest)
		return
	}

	var receivedUser User
	err := json.NewDecoder(r.Body).Decode(&receivedUser)

	if err != nil {
		http.Error(w, "Could not parse incoming body to a user type", http.StatusBadRequest)
		return
	}

	if len(receivedUser.Username) == 0 {
		receivedUser.Username = existingUser.Username
	}

	if len(receivedUser.Password) == 0 {
		receivedUser.Password = existingUser.Password
	} else {
		b := md5.Sum([]byte(receivedUser.Password))
		hashedPass := string(b[:])
		receivedUser.Password = hashedPass
	}
	receivedUser.UID = parsedUuid

	users[parsedUuid] = receivedUser

	responseBytes, _ := json.Marshal(CreateAUserToReturn(receivedUser))
	_, _ = w.Write(responseBytes)
}

func CreateAUserToReturn(u User) UserToReturn {
	return UserToReturn{u.Username, u.UID}
}

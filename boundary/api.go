package boundary

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jamesmawm/golang-user-microservice/control"
	"github.com/jamesmawm/golang-user-microservice/dto"
	"github.com/jamesmawm/golang-user-microservice/model"

	"github.com/google/uuid"
)

type UserApi struct {
	users *control.UserService
}

func NewUserApi() *UserApi {
	return &UserApi{
		users: control.NewUserService(),
	}
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

func (api *UserApi) OnSignup(w http.ResponseWriter, r *http.Request) {
	var dto *dto.UserDto

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(dto.Username) == 0 || len(dto.Password) == 0 {
		http.Error(w, "Invalid params", http.StatusBadRequest)
		return
	}

	var hashedPass string
	var b [16]byte
	b = md5.Sum([]byte(dto.Password))
	hashedPass = string(b[:])

	user := &model.User{
		UID:      uuid.New(),
		Username: dto.Username,
		Password: hashedPass,
	}

	api.users.Create(user)

	w.Header().Set("Location", fmt.Sprintf("/api/users/%s", user.UID.String()))
	w.WriteHeader(http.StatusCreated)
}

func (api *UserApi) OnDeleteUser(w http.ResponseWriter, r *http.Request) {
	ss := strings.Split(r.URL.Path, "/")
	uid := ss[len(ss)-1]

	if len(uid) == 0 {
		http.Error(w, "Invalid query param", http.StatusBadRequest)
		return
	}

	user := api.users.FindOneByUid(uid)

	if user == nil {
		http.Error(w, "User with uuid "+uid+" does not exist", http.StatusBadRequest)
		return
	}

	api.users.Delete(user)

	return
}

func (api *UserApi) OnGetUser(w http.ResponseWriter, r *http.Request) {
	ss := strings.Split(r.URL.Path, "/")
	uid := ss[len(ss)-1]

	if len(uid) == 0 {
		http.Error(w, "Invalid query param", http.StatusBadRequest)
		return
	}

	user := api.users.FindOneByUid(uid)

	if user == nil {
		http.Error(w, "User with uuid "+uid+" does not exist", http.StatusBadRequest)
		return
	}

	responseBytes, _ := json.Marshal(convertToDto(*user))
	_, _ = w.Write(responseBytes)
}

func (api *UserApi) OnUpdateUser(w http.ResponseWriter, r *http.Request) {
	ss := strings.Split(r.URL.Path, "/")
	uid := ss[len(ss)-1]

	if len(uid) == 0 {
		http.Error(w, "Invalid query param", http.StatusBadRequest)
		return
	}

	user := api.users.FindOneByUid(uid)

	if user == nil {
		http.Error(w, "User with uuid "+uid+" does not exist", http.StatusBadRequest)
		return
	}

	var dto dto.UserDto
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, "Could not parse incoming body to a user type", http.StatusBadRequest)
		return
	}

	if len(dto.Username) > 0 {
		user.Username = dto.Username
	}

	if len(dto.Password) > 0 {
		b := md5.Sum([]byte(dto.Password))
		hashedPass := string(b[:])
		user.Password = hashedPass
	}

	api.users.Update(user)

	responseBytes, _ := json.Marshal(convertToDto(*user))
	_, _ = w.Write(responseBytes)
}

func convertToDto(u model.User) dto.UserDto {
	return dto.UserDto{u.Username, "", u.UID}
}

package control

import (
	"errors"

	"github.com/jamesmawm/golang-user-microservice/data"
	"github.com/jamesmawm/golang-user-microservice/model"

	"gorm.io/gorm"
)

type UserService struct {
	db    *gorm.DB
	table string
}

func NewUserService() *UserService {
	return &UserService{
		db:    data.GetDatabase(),
		table: "users",
	}
}

func (service *UserService) Create(entity *model.User) {
	service.db.Create(&entity)
}

func (service *UserService) Delete(entity *model.User) {
	service.db.Delete(&entity)
}

func (service *UserService) FindOneByUid(uid string) *model.User {
	var user model.User

	res := service.db.First(&user, "uid = ?", uid)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &user
}

func (service *UserService) Update(entity *model.User) {
	service.db.Save(&entity)
}

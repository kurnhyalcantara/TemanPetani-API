package service

import (
	"github.com/kurnhyalcantara/TemanPetani-API/apis/users"
	"github.com/kurnhyalcantara/TemanPetani-API/apis/users/model"
)

type UserServiceInterface interface {
	AddUser(user *model.CreateUser) error
	ShowAllUsers(limit int, offset uint) ([]*users.User, error)
	ShowUser(ID uint) (*users.User, error)
	EditUser(ID uint, user *model.UpdateUser) error
	DeleteUser(ID uint)
}

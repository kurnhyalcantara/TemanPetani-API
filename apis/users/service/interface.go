package service

import (
	"github.com/kurnhyalcantara/TemanPetani-API/apis/users"
)

type UserServiceInterface interface {
	RegisterUser(user *users.CreateUser) error
	ShowAllUsers(limit int, offset uint) ([]*users.User, error)
	ShowUser(ID uint) (*users.User, error)
	EditUser(ID uint, user *users.UpdateUser) error
	DeleteUser(ID uint)
}

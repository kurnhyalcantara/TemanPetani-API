package repository

import (
	"github.com/kurnhyalcantara/TemanPetani-API/apis/users"
	"github.com/kurnhyalcantara/TemanPetani-API/apis/users/model"
)

type UserRepoInterface interface {
	Insert(user *model.User) (ID uint, err error)
	SelectAll(limit int, offset uint) ([]*users.User, error)
	Select(ID uint) (*users.User, error)
	Update(ID uint, core *model.User) (*users.User, error)
	Delete(ID uint) error
}

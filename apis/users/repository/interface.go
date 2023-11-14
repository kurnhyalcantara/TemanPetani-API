package repository

import "github.com/kurnhyalcantara/TemanPetani-API/apis/users/model"

type UserRepoInterface interface {
	Create(core *model.CreateUser) error
	GetAll(limit int, offset uint) ([]*model.User, error)
	Get(ID uint) (*model.User, error)
	Update(ID uint, core *model.User) error
	Delete(ID uint)
}


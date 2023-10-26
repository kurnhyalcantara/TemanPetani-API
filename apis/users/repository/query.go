package repository

import (
	"github.com/kurnhyalcantara/TemanPetani-API/apis/users/model"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

// Create implements UserRepoInterface.
func (*userRepo) Create(core *model.CreateUser) error {
	panic("unimplemented")
}

// Delete implements UserRepoInterface.
func (*userRepo) Delete(ID uint) {
	panic("unimplemented")
}

// Get implements UserRepoInterface.
func (*userRepo) Get(ID uint) (*model.User, error) {
	panic("unimplemented")
}

// GetAll implements UserRepoInterface.
func (*userRepo) GetAll(limit int, offset uint) ([]*model.User, error) {
	panic("unimplemented")
}

// Update implements UserRepoInterface.
func (*userRepo) Update(ID uint, core *model.UpdateUser) error {
	panic("unimplemented")
}

func New(db *gorm.DB) UserRepoInterface {
	return &userRepo{
		db: db,
	}
}

package repository

import (
	"github.com/kurnhyalcantara/TemanPetani-API/apis/users/model"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

// Create implements UserRepoInterface
func (repo *userRepo) Create(core *model.CreateUser) error {
	if tx := repo.db.Create(&core); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Get implements UserRepoInterface
func (repo *userRepo) Get(ID uint) (*model.User, error) {
	var user model.User
	if tx := repo.db.First(&user, ID); tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

// GetAll implements UserRepoInterface
func (repo *userRepo) GetAll(limit int, offset uint) ([]*model.User, error) {
	panic("unimplemented")
}

// Delete implements UserRepoInterface
func (repo *userRepo) Delete(ID uint) error {
	if tx := repo.db.Delete(&model.User{}, ID); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Update implements UserRepoInterface
func (*userRepo) Update(ID uint, core *model.User) error {
	panic("unimplemented")
}

func New(db *gorm.DB) UserRepoInterface {
	return &userRepo{
		db: db,
	}
}

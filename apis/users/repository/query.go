package repository

import (
	"errors"

	"github.com/kurnhyalcantara/TemanPetani-API/apis/users"
	"github.com/kurnhyalcantara/TemanPetani-API/apis/users/model"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

// Insert implements UserRepoInterface
func (repo *userRepo) Insert(user *model.User) (ID uint, err error) {
	if tx := repo.db.Create(&user); tx.Error != nil {
		return 0, tx.Error
	}
	return user.ID, nil
}

// Select implements UserRepoInterface
func (repo *userRepo) Select(ID uint) (*users.User, error) {
	var modelUser *model.User
	if tx := repo.db.First(&modelUser, ID); tx.Error != nil {
		return nil, tx.Error
	}
	user := users.ToUser(modelUser)
	return user, nil
}

// SelectAll implements UserRepoInterface
func (repo *userRepo) SelectAll(limit int, offset uint) ([]*users.User, error) {
	var modelUsers []*model.User
	if limit > 0 || offset > 0 {
		if tx := repo.db.Limit(limit).Offset(int(offset)).Find(&modelUsers); tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		if tx := repo.db.Find(&modelUsers); tx.Error != nil {
			return nil, tx.Error
		}
	}
	users := users.ToUsers(modelUsers)
	return users, nil
}

// Update implements UserRepoInterface
func (repo *userRepo) Update(ID uint, user *model.User) (*users.User, error) {
	var modelUser *model.User
	if tx := repo.db.First(&modelUser, ID); tx.Error != nil {
		return nil, tx.Error
	}
	if tx := repo.db.Model(&modelUser).Updates(&user); tx.Error != nil {
		return nil, tx.Error
	}
	var modifiedUser *model.User
	if tx := repo.db.First(&modifiedUser, ID); tx.Error != nil {
		return nil, tx.Error
	}
	userData := users.ToUser(modifiedUser)
	return userData, nil
}

// Delete implements UserRepoInterface
func (repo *userRepo) Delete(ID uint) error {
	tx := repo.db.Delete(&model.User{}, ID)
	if tx.RowsAffected == 0 {
		return errors.New("invalid id, no rows affected")
	}
	return nil
}

func New(db *gorm.DB) UserRepoInterface {
	return &userRepo{
		db: db,
	}
}

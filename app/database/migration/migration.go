package migration

import (
	_userModel "github.com/kurnhyalcantara/TemanPetani-API/apis/users/model"
	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) error {
	return db.AutoMigrate(_userModel.User{})
}

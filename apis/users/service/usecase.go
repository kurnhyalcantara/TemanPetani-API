package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/kurnhyalcantara/TemanPetani-API/apis/users"
	"github.com/kurnhyalcantara/TemanPetani-API/apis/users/model"
	"github.com/kurnhyalcantara/TemanPetani-API/apis/users/repository"
	"github.com/kurnhyalcantara/TemanPetani-API/app/libs"
)

type UserService struct {
	userRepo  repository.UserRepoInterface
	validator *validator.Validate
	hashing   libs.Bcrypt
}

// AddUser implements UserServiceInterface
func (service *UserService) AddUser(user *model.CreateUser) error {
	if errValidate := service.validator.Struct(user); errValidate != nil {
		return errValidate
	}

	hashedPassword, errHash := service.hashing.HashPassword(user.Password)
	if errHash != nil {
		return errHash
	}

	user.Password = string(hashedPassword)

	if errInsert := service.userRepo.Create(user); errInsert != nil {
		return errInsert
	}

	return nil
}

// DeleteUser implements UserServiceInterface
func (*UserService) DeleteUser(ID uint) {
	panic("unimplemented")
}

// EditUser implements UserServiceInterface
func (*UserService) EditUser(ID uint, user *model.UpdateUser) error {
	panic("unimplemented")
}

// ShowAllUsers implements UserServiceInterface
func (*UserService) ShowAllUsers(limit int, offset uint) ([]*users.User, error) {
	panic("unimplemented")
}

// ShowUser implements UserServiceInterface
func (*UserService) ShowUser(ID uint) (*users.User, error) {
	panic("unimplemented")
}

func New(userRepo repository.UserRepoInterface) UserServiceInterface {
	return &UserService{
		userRepo:  userRepo,
		validator: validator.New(),
		hashing:   *libs.NewBcrypt(),
	}
}

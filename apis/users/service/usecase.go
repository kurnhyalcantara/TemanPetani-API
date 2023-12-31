package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/kurnhyalcantara/TemanPetani-API/apis/users"
	"github.com/kurnhyalcantara/TemanPetani-API/apis/users/repository"
	"github.com/kurnhyalcantara/TemanPetani-API/app/libs"
)

type UserService struct {
	userRepo  repository.UserRepoInterface
	validator *validator.Validate
	hashing   libs.Bcrypt
}

// RegisterUser implements UserServiceInterface
func (*UserService) RegisterUser(user *users.CreateUser) error {
	panic("unimplemented")
}

// DeleteUser implements UserServiceInterface
func (*UserService) DeleteUser(ID uint) {
	panic("unimplemented")
}

// EditUser implements UserServiceInterface
func (*UserService) EditUser(ID uint, user *users.UpdateUser) error {
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
	v := validator.New()
	// Register the custom validation function
	if err := v.RegisterValidation("validatePassword", users.ValidatePassword); err != nil {
		panic("Failed to register validation function: " + err.Error())
	}

	return &UserService{
		userRepo:  userRepo,
		validator: v,
		hashing:   *libs.NewBcrypt(),
	}
}

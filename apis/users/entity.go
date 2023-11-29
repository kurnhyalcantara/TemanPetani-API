package users

import (
	"time"
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/kurnhyalcantara/TemanPetani-API/apis/users/model"
)

type User struct {
	ID            uint      `json:"id"`
	FullName      string    `json:"fullname"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Role          string    `json:"role"`
	Address       string    `json:"address,omitempty"`
	Avatar        string    `json:"avatar,omitempty"`
	Bank          string    `json:"bank,omitempty"`
	AccountNumber string    `json:"accountNumber,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type CreateUser struct {
	FullName string `json:"fullname" validate:"required,lte=200,gte=5"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,lte=13"`
	Password string `json:"password" validate:"required,gte=8,validatePassword"`
	Address  string `json:"address"`
}

type UpdateUser struct {
	FullName      string `json:"fullname" validate:"required,lte=200,gte=5"`
	Address       string `json:"address"`
	Avatar        string `json:"avatar"`
	Bank          string `json:"bank"`
	AccountNumber string `json:"accountNumber"`
}

func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	var lower, upper, number, special bool
	for _, c := range password {
		switch {
		case unicode.IsLower(c):
			lower = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsNumber(c):
			number = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}
	}

	return lower && upper && number && special
}

func ToUser(user *model.User) *User {
	return &User{
		ID:            user.ID,
		FullName:      user.FullName,
		Email:         user.Email,
		Phone:         user.Phone,
		Role:          user.Role,
		Address:       user.Address,
		Avatar:        user.Avatar,
		Bank:          user.Bank,
		AccountNumber: user.AccountNumber,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}

func ToUsers(users []*model.User) []*User {
	result := make([]*User, len(users))
	for i, user := range users {
		result[i] = ToUser(user)
	}
	return result
}

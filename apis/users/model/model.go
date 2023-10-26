package model

import (
	"time"
	"unicode"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	ID            uint64 `gorm:"primaryKey"`
	FullName      string `gorm:"notNull"`
	Email         string `gorm:"unique;notNull"`
	Phone         string `gorm:"unique;notNull"`
	Password      string `gorm:"notNull"`
	Role          string `gorm:"type:enum('admin', 'user');default:'user'"`
	Address       string `gorm:"type:text"`
	Avatar        string
	Bank          string
	AccountNumber string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
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

func validatePassword(fl validator.FieldLevel) bool {
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

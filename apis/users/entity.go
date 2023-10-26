package users

import (
	"time"

	"github.com/kurnhyalcantara/TemanPetani-API/apis/users/model"
)

type User struct {
	ID            uint64    `json:"id"`
	FullName      string    `json:"fullname"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Role          string    `json:"role"`
	Address       string    `json:"address"`
	Avatar        string    `json:"avatar"`
	Bank          string    `json:"bank"`
	AccountNumber string    `json:"accountNumber"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func ToUser(user *model.User) *User {
	return &User{
		ID:            user.ID,
		FullName:      user.FullName,
		Email:         user.Email,
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

package service

import (
	"errors"
	"testing"

	"github.com/kurnhyalcantara/TemanPetani-API/apis/users/model"
	"github.com/kurnhyalcantara/TemanPetani-API/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddUser(t *testing.T) {
	repo := new(mocks.UserRepo)
	srv := New(repo)

	t.Run("[Success] Add User with Valid Payload", func(t *testing.T) {
		newUser := model.CreateUser{
			FullName: "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "1234567890123",
			Password: "SecurePwd@123",
			Address:  "123 Main Street, Cityville",
		}
		repo.On("Create", mock.Anything).Return(nil).Once()
		err := srv.AddUser(&newUser)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("[Fail] Add User with Invalid Payload", func(t *testing.T) {
		repo.On("Create", mock.Anything).Return(errors.New("error validator: ")).Once()
		err := srv.AddUser(&model.CreateUser{})
		assert.NotNil(t, err)
	})

	t.Run("[Fail] Add User with Error", func(t *testing.T) {
		newUser := model.CreateUser{
			FullName: "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "1234567890123",
			Password: "SecurePwd@123",
			Address:  "123 Main Street, Cityville",
		}
		repo.On("Create", mock.Anything).Return(errors.New("")).Once()
		err := srv.AddUser(&newUser)
		assert.Error(t, err)
	})
}

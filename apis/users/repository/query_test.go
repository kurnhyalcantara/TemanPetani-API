package repository

import (
	"testing"

	"github.com/kurnhyalcantara/TemanPetani-API/apis/users"
	"github.com/kurnhyalcantara/TemanPetani-API/apis/users/model"
	mysqldb "github.com/kurnhyalcantara/TemanPetani-API/app/database/mysql"
	"github.com/stretchr/testify/assert"
)

var db = mysqldb.InitDB()
var currentId uint

func TestInsertUser(t *testing.T) {
	userRepo := New(db)
	t.Run("Case #1: Success Insert User", func(t *testing.T) {
		db.Migrator().DropTable(&model.User{})
		db.AutoMigrate(&model.User{})
		mockUser := &model.User{
			FullName: "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "1234567890123",
			Password: "strongpwd",
			Address:  "123 Main Street, Cityville",
		}

		userId, err := userRepo.Insert(mockUser)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), userId)
		currentId = userId
	})

	t.Run("Case #2: Failed Insert User - Duplicate Email", func(t *testing.T) {
		existingUser := &model.User{
			FullName: "Jane Doe",
			Email:    "jane.doe@example.com",
			Phone:    "9876543210987",
			Password: "securepwd",
			Address:  "456 Oak Street, Townsville",
		}
		_, err := userRepo.Insert(existingUser)
		assert.Nil(t, err)

		duplicateUser := &model.User{
			FullName: "Duplicate User",
			Email:    "jane.doe@example.com", // Email yang sudah terdaftar
			Phone:    "5555555555555",
			Password: "duplicatepwd",
			Address:  "789 Pine Street, Villagetown",
		}
		_, err = userRepo.Insert(duplicateUser)
		assert.NotNil(t, err)
		assert.Error(t, err, "Error should be returned for duplicate email")
	})
}

func TestSelectUser(t *testing.T) {
	userRepo := New(db)

	t.Run("Case #1: Success Select User", func(t *testing.T) {
		mockUser := &users.User{
			ID:       currentId,
			FullName: "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "1234567890123",
			Role:     "user",
			Address:  "123 Main Street, Cityville",
		}
		user, err := userRepo.Select(currentId)
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Email, user.Email)
		assert.Equal(t, mockUser.Phone, user.Phone)
	})

	t.Run("Case #2: Failed Select User - Invalid ID", func(t *testing.T) {
		_, err := userRepo.Select(0)
		assert.NotNil(t, err)
	})
}

func TestSelectAllUser(t *testing.T) {
	userRepo := New(db)

	t.Run("Case #1: Success Select All Users", func(t *testing.T) {
		users, err := userRepo.SelectAll(0, 0)
		assert.Nil(t, err)
		assert.Len(t, users, 2)
	})

	t.Run("Case #2: Success Select Some Users with Limit and Offset", func(t *testing.T) {
		existingUser := &model.User{
			FullName: "Jane Doe",
			Email:    "jane.doe@example.com",
			Phone:    "9876543210987",
			Password: "securepwd",
			Address:  "456 Oak Street, Townsville",
		}
		users, err := userRepo.SelectAll(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, existingUser.Email, users[0].Email)
	})

	t.Run("Case #3: Failed Select Some Users with Invalid Limit and Offset", func(t *testing.T) {
		_, err := userRepo.SelectAll(-1, 3)
		assert.Error(t, err, "Select all offset limit error")
	})
}

func TestUpdateUser(t *testing.T) {
	userRepo := New(db)
	mockUser := &model.User{
		FullName: "John Doe modified",
		Email:    "john.doe@example.com",
		Phone:    "1234567890123",
		Password: "strongpwd",
		Address:  "123 Main Street, Cityville",
	}
	t.Run("Case #1: Success Updated Data", func(t *testing.T) {
		userModified, err := userRepo.Update(currentId, mockUser)
		if assert.NoError(t, err) {
			assert.Equal(t, mockUser.FullName, userModified.FullName)
		}
	})

	t.Run("Case #2: Failed Select User Data with Invalid ID", func(t *testing.T) {
		userModified, err := userRepo.Update(0, mockUser)
		if assert.NotNil(t, err) {
			assert.Error(t, err, "Error get user data with invalid id")
		}
		assert.Nil(t, userModified)
	})
}

func TestDeleteUser(t *testing.T) {
	userRepo := New(db)

	t.Run("Case #1: Success Deleted Data", func(t *testing.T) {
		err := userRepo.Delete(currentId)
		assert.Nil(t, err)

		user, err := userRepo.Select(currentId)
		assert.NotNil(t, err)
		assert.Nil(t, user)
	})

	t.Run("Case #2: Failed Delete Data with Invalid ID", func(t *testing.T) {
		err := userRepo.Delete(uint(999))
		assert.Error(t, err, "Delete data with invalid ID")
	})
}

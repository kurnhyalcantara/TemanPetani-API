package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint   `gorm:"primaryKey"`
	FullName      string `gorm:"not null;type:varchar(100)"`
	Email         string `gorm:"unique;not null"`
	Phone         string `gorm:"unique;not null"`
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

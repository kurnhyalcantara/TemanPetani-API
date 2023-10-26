package libs

import (
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct{}

func NewBcrypt() *Bcrypt {
	return &Bcrypt{}
}

func (b *Bcrypt) HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (b *Bcrypt) ComparePassword(hashedPassowrd, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassowrd), []byte(password))
}

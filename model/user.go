
package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct{
	Id int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"` 
}

func (User) TableName() string {
	return "User"
}

type Error struct{
	Message string `json:"message"` 
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
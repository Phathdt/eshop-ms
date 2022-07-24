package models

import (
	"golang.org/x/crypto/bcrypt"
	"user_api/pkg/sdkcm"
)

type User struct {
	sdkcm.SQLModel `json:",inline"`
	Email          string `json:"email"`
	Password       string `json:"password"`
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (User) TableName() string {
	return "users"
}

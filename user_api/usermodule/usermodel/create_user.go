package models

import (
	"golang.org/x/crypto/bcrypt"
	"user_api/pkg/sdkcm"
)

type CreateUser struct {
	sdkcm.SQLModel `json:",inline"`
	Email          string `json:"email" validate:"required"`
	Password       string `json:"password" validate:"required"`
}

func (user *CreateUser) SetPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

	return string(hashedPassword)
}

func (CreateUser) TableName() string {
	return User{}.TableName()
}

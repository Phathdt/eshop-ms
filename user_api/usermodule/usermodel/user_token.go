package models

import (
	"user_api/pkg/sdkcm"
)

type UserToken struct {
	sdkcm.SQLModel `json:",inline"`
	UserID         uint32 `json:"user_id"`
	Token          string `json:"token"`
}

func (UserToken) TableName() string {
	return "user_tokens"
}

package models

import (
	"user_api/pkg/sdkcm"
)

type UserToken struct {
	sdkcm.SQLModel `json:",inline"`
	UserID         uint32          `json:"user_id"`
	Token          string          `json:"token"`
	ExpiredAt      *sdkcm.JSONTime `json:"expired_at"`
}

func (UserToken) TableName() string {
	return "user_tokens"
}

package userhandler

import (
	"context"

	models "user_api/usermodule/usermodel"
)

type LoginUserRepo interface {
	FindUser(ctx context.Context, data map[string]interface{}) (*models.User, error)
	CreateUserToken(ctx context.Context, user *models.User) (*string, error)
}

type loginUserHdl struct {
	repo LoginUserRepo
}

func NewLoginUserHdl(store LoginUserRepo) *loginUserHdl {
	return &loginUserHdl{repo: store}
}

func (h *loginUserHdl) Response(ctx context.Context, data *models.LoginUser) (*string, error) {
	user, err := h.repo.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, err
	}

	if err = user.ComparePassword(data.Password); err != nil {
		return nil, err
	}

	token, err := h.repo.CreateUserToken(ctx, user)
	if err != nil {
		return nil, err
	}

	return token, nil
}

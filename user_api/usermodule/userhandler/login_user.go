package userhandler

import (
	"context"

	"user_api/middleware"
	models "user_api/usermodule/usermodel"
)

type LoginUserRepo interface {
	FindUser(ctx context.Context, data map[string]interface{}) (*models.User, error)
	CreateUserToken(ctx context.Context, userID uint32, token string) error
}

type loginUserHdl struct {
	store LoginUserRepo
}

func NewLoginUserHdl(store LoginUserRepo) *loginUserHdl {
	return &loginUserHdl{store: store}
}

func (h *loginUserHdl) Response(ctx context.Context, data *models.LoginUser) (*string, error) {
	user, err := h.store.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, err
	}

	if err = user.ComparePassword(data.Password); err != nil {
		return nil, err
	}

	token, err := middleware.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	if err = h.store.CreateUserToken(ctx, user.ID, token); err != nil {
		return nil, err
	}

	return &token, nil
}

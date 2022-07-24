package userhandler

import (
	"context"

	models "user_api/usermodule/usermodel"
)

type CreateUserRepo interface {
	CreateUser(ctx context.Context, data *models.CreateUser) (uint32, error)
}

type createUserHdl struct {
	repo CreateUserRepo
}

func NewCreateUserHdl(repo CreateUserRepo) *createUserHdl {
	return &createUserHdl{repo: repo}
}

func (h *createUserHdl) Response(ctx context.Context, data *models.CreateUser) error {
	_, err := h.repo.CreateUser(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

package userrepo

import (
	"context"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"user_api/pkg/sdkcm"
	models "user_api/usermodule/usermodel"
)

type UserStorage interface {
	GetUserByCondition(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*models.User, error)
	CreateUser(ctx context.Context, data *models.CreateUser) (uint32, error)
	CreateUserToken(ctx context.Context, data *models.UserToken) error
}

type repo struct {
	store  UserStorage
	logger *logrus.Logger
}

func NewRepo(store UserStorage, logger *logrus.Logger) *repo {
	return &repo{store: store, logger: logger}
}

func (r *repo) CreateUser(ctx context.Context, data *models.CreateUser) (uint32, error) {
	user := models.CreateUser{
		SQLModel: *sdkcm.NewSQLModel(),
		Email:    data.Email,
		Password: data.SetPassword(data.Password),
	}

	return r.store.CreateUser(ctx, &user)
}

func (r *repo) FindUser(ctx context.Context, data map[string]interface{}) (*models.User, error) {
	return r.store.GetUserByCondition(ctx, data)
}

func (r *repo) CreateUserToken(ctx context.Context, user *models.User) (*string, error) {
	claims := jwt.MapClaims{
		"email": user.Email,
		"id":    user.ID,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	signs := strings.Split(t, ".")

	userToken := models.UserToken{
		SQLModel: *sdkcm.NewSQLModel(),
		UserID:   user.ID,
		Token:    signs[2],
	}

	if err = r.store.CreateUserToken(ctx, &userToken); err != nil {
		return nil, err
	}

	return &t, nil
}

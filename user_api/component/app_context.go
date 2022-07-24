package component

import (
	"context"
	"fmt"
	"sync"
	"time"

	"user_api/pkg/config"
	"user_api/pkg/logger"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AppContext struct {
	DB       *gorm.DB
	RdClient *redis.Client
	Logger   *logrus.Logger
}

func NewAppContext(ctx context.Context) (*AppContext, error) {
	cfg := config.Config

	l := logger.New(cfg.App.LogLevel)

	db, err := NewGormService()
	if err != nil {
		return nil, err
	}

	rdClient, err := NewRedisService(ctx)
	if err != nil {
		return nil, err
	}

	return &AppContext{
		DB:       db,
		RdClient: rdClient,
		Logger:   l,
	}, nil
}

func (appContext *AppContext) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	var errs []error

	go func() {
		defer wg.Done()
		if appContext.DB != nil {
			db, _ := appContext.DB.DB()

			err := db.Close()
			if err != nil {
				errs = append(errs, err)
			}
		}
	}()

	go func() {
		defer wg.Done()
		if appContext.RdClient != nil {
			if err := appContext.RdClient.Close(); err != nil {
				errs = append(errs, err)
			}
		}
	}()

	wg.Wait()

	var closeErr error
	for _, err := range errs {
		if closeErr == nil {
			closeErr = err
		} else {
			closeErr = fmt.Errorf("%v | %v", closeErr, err)
		}
	}

	if closeErr != nil {
		return closeErr
	}

	fmt.Println("Shutdown all service successfully")

	return ctx.Err()
}

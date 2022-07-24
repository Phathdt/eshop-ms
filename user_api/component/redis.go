package component

import (
	"context"
	"fmt"

	"user_api/pkg/config"

	"github.com/go-redis/redis/v8"
)

func NewRedisService(ctx context.Context) (*redis.Client, error) {
	cfg := config.Config
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.REDIS.Host,
		DB:   cfg.REDIS.DB,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("rdb.Ping %w", err)
	}

	return rdb, nil
}

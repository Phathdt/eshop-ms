package config

import (
	"github.com/caarlos0/env/v6"
)

var Config = FromEnv()

type config struct {
	App struct {
		Environment string `env:"USER_API_APP_ENV"   envDefault:"dev"`
		LogLevel    string `env:"USER_API_LOG_LEVEL" envDefault:"INFO"`
	}
	HTTP struct {
		Port int `env:"USER_API_HTTP_PORT" envDefault:"4000"`
	}
	GORM struct {
		JsonLogger bool `env:"USER_API_GORM_JSON_LOGGER" envDefault:"false"`
	}
	POSTGRES struct {
		Host     string `env:"USER_API_POSTGRES_HOST"     envDefault:"0.0.0.0"`
		Port     int    `env:"USER_API_POSTGRES_PORT"     envDefault:"5432"`
		User     string `env:"USER_API_POSTGRES_USER"     envDefault:"postgres"`
		Pass     string `env:"USER_API_POSTGRES_PASS"     envDefault:"123123123"`
		Database string `env:"USER_API_POSTGRES_DATABASE" envDefault:"user_api"`
		Sslmode  string `env:"USER_API_POSTGRES_SSLMODE"  envDefault:"disable"`
	}
	REDIS struct {
		Host string `env:"USER_API_REDIS_HOST" envDefault:"localhost:6379"`
		DB   int    `env:"USER_API_REDIS_DB"   envDefault:"0"`
	}
}

func FromEnv() *config {
	var c config

	if err := env.Parse(&c.App); err != nil {
		panic(err)
	}
	if err := env.Parse(&c.HTTP); err != nil {
		panic(err)
	}
	if err := env.Parse(&c.GORM); err != nil {
		panic(err)
	}
	if err := env.Parse(&c.POSTGRES); err != nil {
		panic(err)
	}
	if err := env.Parse(&c.REDIS); err != nil {
		panic(err)
	}

	return &c
}

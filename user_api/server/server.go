package server

import (
	"fmt"

	jwtware "github.com/gofiber/jwt/v3"
	"user_api/component"
	"user_api/middleware"
	"user_api/pkg/config"
	"user_api/usermodule/usertransport/userfiber"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type server struct {
	AppContext *component.AppContext
	app        *fiber.App
}

func NewServer(appContext *component.AppContext) *server {
	return &server{AppContext: appContext, app: fiber.New()}
}

func ping() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(&fiber.Map{
			"msg": "pong",
		})
	}
}

func (s *server) Run() error {
	app := s.app
	cfg := config.Config

	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(cors.New())

	app.Use(middleware.Recover(s.AppContext.Logger))

	app.Get("/", ping())
	app.Get("/ping", ping())

	app.Post("/api/users", userfiber.CreateUser(s.AppContext))
	app.Post("/api/users/login", userfiber.Login(s.AppContext))

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	addr := fmt.Sprintf(":%d", cfg.HTTP.Port)
	if err := app.Listen(addr); err != nil {
		return err
	}
	return nil
}

func (s *server) Shutdown() {
	_ = s.app.Shutdown()
}

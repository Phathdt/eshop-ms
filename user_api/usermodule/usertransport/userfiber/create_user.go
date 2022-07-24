package userfiber

import (
	"net/http"

	"user_api/component"
	"user_api/pkg/sdkcm"
	"user_api/pkg/validation"
	"user_api/usermodule/userhandler"
	models "user_api/usermodule/usermodel"
	"user_api/usermodule/userrepo"
	"user_api/usermodule/userstorage"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(appContext *component.AppContext) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var p models.CreateUser

		if err := ctx.BodyParser(&p); err != nil {
			panic(err)
		}

		if err := validation.Validate(p); err != nil {
			panic(err)
		}

		storage := userstorage.NewSqlStorage(appContext.DB)
		repo := userrepo.NewRepo(storage, appContext.Logger)
		hdl := userhandler.NewCreateUserHdl(repo)

		err := hdl.Response(ctx.Context(), &p)
		if err != nil {
			panic(err)
		}

		return ctx.Status(http.StatusOK).JSON(sdkcm.SimpleSuccessResponse(map[string]interface{}{
			"msg": "ok",
		}))
	}
}

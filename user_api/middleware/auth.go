package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"user_api/component"
	"user_api/usermodule/userrepo"
	"user_api/usermodule/userstorage"
)

func IsAuthenticated(appContext *component.AppContext) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Locals("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		storage := userstorage.NewSqlStorage(appContext.DB)
		repo := userrepo.NewRepo(storage, appContext.Logger)
		user, err := repo.FindUser(ctx.Context(), map[string]interface{}{"id": claims["id"].(float64)})
		if err != nil {
			return ctx.JSON(fiber.Map{
				"message": "unauthenticated",
			})
		}

		userToken, err := repo.FindUserToken(ctx.Context(), map[string]interface{}{
			"user_id": user.ID,
			"token":   claims["key"].(string),
		})

		if err != nil {
			return ctx.JSON(fiber.Map{
				"message": "unauthenticated",
			})
		}

		if !userToken.Active {
			return ctx.JSON(fiber.Map{
				"message": "unauthenticated",
			})
		}

		ctx.Context().SetUserValue("user", user)

		return ctx.Next()
	}
}

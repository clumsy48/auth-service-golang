package middleware

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
	"webserser/domain"
	"webserser/models"
)

func NewAuthMiddleware(as domain.AuthService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")
		token = strings.Replace(token, "bearer ", "", 1)
		userId, err := as.Validate(ctx.Context(), domain.AuthTokenID(token))
		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(
				models.Response[any]{
					Success: false,
					Message: err.Error(),
				},
			)
		}
		customHeaders := make(map[string]string)
		customHeaders["user_id"] = userId
		ctx.Locals("customHeaders", customHeaders)
		return ctx.Next()
	}
}

package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"webserser/domain"
	"webserser/models"
)

type AuthHandler struct {
	as domain.AuthService
}

func NewAuthHandler(authService domain.AuthService) AuthHandler {
	return AuthHandler{as: authService}
}

func (ah AuthHandler) Logout(c *fiber.Ctx) error {

	customHeaders := c.Locals("customHeaders")
	userId := customHeaders.(map[string]string)["user_id"]
	err := ah.as.Expire(c.Context(), userId)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(models.Response[string]{
		Success: true,
	})
}

func (ah AuthHandler) Login(c *fiber.Ctx) error {
	var authParams domain.AuthParams

	if err := c.BodyParser(&authParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	tokenID, err := ah.as.Authenticate(c.Context(), authParams)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(models.Response[string]{
		Success: true,
		Message: fmt.Sprintf("%s", tokenID),
	})
}

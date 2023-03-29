package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"webserser/domain"
	"webserser/models"
)

type OnboardingHandler struct {
	os domain.OnboardingService
}

func NewOnboardingHandler(onboardingService domain.OnboardingService) OnboardingHandler {
	return OnboardingHandler{os: onboardingService}
}

func (oh OnboardingHandler) SignUp(c *fiber.Ctx) error {
	var user domain.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	err := oh.os.SignUp(c.Context(), user)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(models.Response[string]{
		Success: true,
		Message: "sign up successful",
	})
}

func (oh OnboardingHandler) Delete(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

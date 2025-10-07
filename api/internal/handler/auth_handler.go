package handler

import (
	"bri-edc/api/internal/services"
	"bri-edc/api/internal/validations"
	"bri-edc/api/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	s *services.AuthService
}

func NewAuthHandler(s *services.AuthService) *AuthHandler {
	return &AuthHandler{s: s}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	req, err, code := utils.ValidateAndBind[validations.AuthRequest](c)
	if err != nil {
		return c.Status(code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	token, err := h.s.Login(*req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}

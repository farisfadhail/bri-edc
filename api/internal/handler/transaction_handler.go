package handler

import (
	"bri-edc/api/internal/services"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	s *services.TransactionService
}

func NewTransactionHandler(s *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{s: s}
}

func (h *TransactionHandler) Sale(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"rc":      "SUCCESS",
		"message": "Sale transaction OK 8081",
	})
}

func (h *TransactionHandler) Settlement(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"rc":      "SUCCESS",
		"message": "Settlement transaction OK",
	})
}

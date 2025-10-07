package handler

import (
	"bri-edc/api/internal/services"
	"bri-edc/api/internal/validations"
	"bri-edc/api/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	s *services.TransactionService
}

func NewTransactionHandler(s *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{s: s}
}

func (h *TransactionHandler) Sale(c *fiber.Ctx) error {
	req, err, code := utils.ValidateAndBind[validations.CreateTransactionRequest](c)
	if err != nil {
		return utils.ResponseFailed(c, code, err.Error())

	}

	res, err := h.s.Sale(*req)
	if err != nil {
		return utils.ResponseFailed(c, http.StatusUnprocessableEntity, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *TransactionHandler) Settlement(c *fiber.Ctx) error {
	settlement, err := h.s.Settlement()
	if err != nil {
		return utils.ResponseFailed(c, http.StatusUnprocessableEntity, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(settlement)
}

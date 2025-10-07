package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type APIResponse struct {
	RC           string      `json:"rc"`
	Message      string      `json:"message"`
	Timestamp    time.Time   `json:"timestamp"`
	LocationTime string      `json:"location_time"`
	Payload      interface{} `json:"payload,omitempty"`
}

type PayloadData struct {
	Data interface{} `json:"data"`
}

func ResponseSuccess(c *fiber.Ctx, message string, data interface{}) error {
	loc := time.Now().Location()

	return c.Status(fiber.StatusOK).JSON(APIResponse{
		RC:           "SUCCESS",
		Message:      message,
		Timestamp:    time.Now(),
		LocationTime: loc.String(),
		Payload:      PayloadData{Data: data},
	})
}

func ResponseFailed(c *fiber.Ctx, code int, message string) error {
	loc := time.Now().Location()

	return c.Status(code).JSON(APIResponse{
		RC:           "FAILED",
		Message:      message,
		Timestamp:    time.Now(),
		LocationTime: loc.String(),
		Payload:      PayloadData{Data: nil},
	})
}

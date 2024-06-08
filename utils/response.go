// utils/response.go

package utils

import (
	"github.com/gofiber/fiber/v2"
)

func JSONResponse(c *fiber.Ctx, statusCode int, ok bool, message string, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"ok":      ok,
		"message": message,
		"data":    data,
	})
}

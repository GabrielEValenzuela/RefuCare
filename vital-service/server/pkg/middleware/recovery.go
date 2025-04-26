package middleware

import (
	"github.com/GabrielEValenzuela/RefuCare/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func Recovery() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				logger.Log.Errorw("panic recovered", "error", r)
				c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
		}()
		return c.Next()
	}
}

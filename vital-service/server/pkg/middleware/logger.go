package middleware

import (
	"time"

	"github.com/GabrielEValenzuela/RefuCare/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		stop := time.Since(start)

		logger.Log.Infow("request",
			"method", c.Method(),
			"path", c.Path(),
			"status", c.Response().StatusCode(),
			"duration", stop,
		)

		return err
	}
}

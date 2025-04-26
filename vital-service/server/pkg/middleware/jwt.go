package middleware

import (
	"github.com/GabrielEValenzuela/RefuCare/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" || len(auth) < 8 || auth[:7] != "Bearer " {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		tokenStr := auth[7:]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		c.Locals("user", claims)
		return c.Next()
	}
}

package middleware

import (
	"peluang-server/domain"
	"peluang-server/dto"
	"peluang-server/internal/util"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(
				&dto.HttpResponse{
					Message: domain.ErrEmptyToken.Error(),
					Code:    fiber.StatusUnauthorized,
					Data:    []string{},
				},
			)
		}
		if parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(
				&dto.HttpResponse{
					Message: domain.ErrNoBerearToken.Error(),
					Code:    fiber.StatusUnauthorized,
					Data:    []string{},
				},
			)
		}

		tokenString = parts[1]

		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(
				&dto.HttpResponse{
					Message: domain.ErrEmptyToken.Error(),
					Code:    fiber.StatusUnauthorized,
					Data:    []string{},
				},
			)
		}
		err := util.ValidateToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(
				&dto.HttpResponse{
					Message: domain.ErrInvalidToken.Error(),
					Code:    fiber.StatusUnauthorized,
					Data:    []string{},
				},
			)
		}
		return c.Next()
	}
}

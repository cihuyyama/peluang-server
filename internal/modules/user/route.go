package user

import (
	"peluang-server/domain"
	"peluang-server/dto"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type route struct {
	userService domain.UserService
}

func NewRoute(app *fiber.App, userService domain.UserService) {
	route := route{
		userService,
	}

	api := app.Group("/api")

	api.Post("/auth/register", route.UserRegister)
}
func (r *route) UserRegister(c *fiber.Ctx) error {
	user := new(dto.AuthRequest)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing body",
		})
	}

	if err := validator.New().Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userModel := new(domain.User)
	userModel.Email = user.Email
	userModel.Password = user.Password
	userModel.Username = user.Username
	userModel.Telp = user.Telp

	otp, err := r.userService.Register(userModel, c.Context())
	if err != nil {
		if err == domain.ErrEmailExist {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": domain.ErrEmailExist.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"OTP": otp,
	})
}

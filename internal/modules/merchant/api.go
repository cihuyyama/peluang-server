package merchant

import (
	"peluang-server/domain"
	"peluang-server/dto"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type merchantRoute struct {
	merchatService domain.MerchantService
}

func NewRoute(app *fiber.App, merchatService domain.MerchantService) {
	merchantRoute := &merchantRoute{
		merchatService,
	}

	api := app.Group("/api/v1/merchant")
	{
		api.Get("", merchantRoute.GetAllMerchants)
		api.Get("/:slug", merchantRoute.GetMerchant)
		api.Post("", merchantRoute.CreateMerchant)
		api.Put("/:id", merchantRoute.UpdateMerchant)
		api.Delete("/:id", merchantRoute.DeleteMerchant)
		api.Put("/:id/avatar", merchantRoute.UpdateAvatar)
		api.Post("/:id/images", merchantRoute.CreateImage)
		api.Delete("/:id/images/:image_id", merchantRoute.DeleteImage)
	}
}

func (mr *merchantRoute) CreateMerchant(c *fiber.Ctx) error {
	merchant := new(dto.MerchantRequest)
	if err := c.BodyParser(merchant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing body",
			"code":    fiber.StatusBadRequest,
			"data":    []string{},
		})
	}

	if err := validator.New().Struct(merchant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error validating body",
			"code":    fiber.StatusBadRequest,
			"data":    []string{},
		})
	}

	if err := mr.merchatService.CreateMerchant(merchant); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusInternalServerError,
			"data":    []string{},
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"code":    fiber.StatusCreated,
		"data":    []string{},
	})
}

func (mr *merchantRoute) GetMerchant(c *fiber.Ctx) error {
	merchantSlug := c.Params("slug")
	merchant, err := mr.merchatService.GetMerchant(merchantSlug)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusNotFound,
			"data":    []string{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"code":    fiber.StatusOK,
		"data":    merchant,
	})
}

func (mr *merchantRoute) GetAllMerchants(c *fiber.Ctx) error {
	merchants, err := mr.merchatService.GetAllMerchants()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusInternalServerError,
			"data":    []string{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"code":    fiber.StatusOK,
		"data":    merchants,
	})
}

func (mr *merchantRoute) UpdateMerchant(c *fiber.Ctx) error {
	merchantID := c.Params("id")
	merchant := new(dto.MerchantRequest)
	if err := c.BodyParser(merchant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing body",
			"code":    fiber.StatusBadRequest,
			"data":    []string{},
		})
	}

	if err := validator.New().Struct(merchant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error validating body",
			"code":    fiber.StatusBadRequest,
			"data":    []string{},
		})
	}

	if err := mr.merchatService.UpdateMerchant(merchantID, merchant); err != nil {
		if err == domain.ErrMerchantNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
				"code":    fiber.StatusNotFound,
				"data":    []string{},
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusInternalServerError,
			"data":    []string{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"code":    fiber.StatusOK,
		"data":    []string{},
	})
}

func (mr *merchantRoute) DeleteMerchant(c *fiber.Ctx) error {
	merchantID := c.Params("id")
	if err := mr.merchatService.DeleteMerchant(merchantID); err != nil {
		if err == domain.ErrMerchantNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
				"code":    fiber.StatusNotFound,
				"data":    []string{},
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusInternalServerError,
			"data":    []string{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"code":    fiber.StatusOK,
		"data":    []string{},
	})
}

func (mr *merchantRoute) UpdateAvatar(c *fiber.Ctx) error {
	merchantID := c.Params("id")
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing file",
			"code":    fiber.StatusBadRequest,
			"data":    []string{},
		})
	}

	if err := mr.merchatService.UpdateAvatar(merchantID, file); err != nil {
		if err == domain.ErrMerchantNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
				"code":    fiber.StatusNotFound,
				"data":    []string{},
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusInternalServerError,
			"data":    []string{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"code":    fiber.StatusOK,
		"data":    []string{},
	})
}

func (mr *merchantRoute) CreateImage(c *fiber.Ctx) error {
	merchantID := c.Params("id")
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing file",
			"code":    fiber.StatusBadRequest,
			"data":    []string{},
		})
	}

	if err := mr.merchatService.CreateImage(merchantID, file); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusInternalServerError,
			"data":    []string{},
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"code":    fiber.StatusCreated,
		"data":    []string{},
	})
}

func (mr *merchantRoute) DeleteImage(c *fiber.Ctx) error {
	merchantID := c.Params("id")
	imageID := c.Params("image_id")
	if err := mr.merchatService.DeleteImage(merchantID, imageID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusInternalServerError,
			"data":    []string{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"code":    fiber.StatusOK,
		"data":    []string{},
	})
}

package banner

import (
	"peluang-server/domain"
	"peluang-server/dto"
	"peluang-server/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type bannerRouter struct {
	bannerService domain.BannerService
}

func NewRoute(app *fiber.App, bannerService domain.BannerService) {
	bannerRoute := &bannerRouter{
		bannerService: bannerService,
	}

	api := app.Group("/api/v1/banner")
	{
		api.Get("", bannerRoute.GetAllBanners)
		api.Get("/:id", bannerRoute.GetBanner)
	}

	protected := app.Group("/api/v1/banner", middleware.Authenticate())
	{
		protected.Post("", bannerRoute.CreateBanner)
		protected.Delete("/:id", bannerRoute.DeleteBanner)
	}
}

func (br *bannerRouter) CreateBanner(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing file",
			"code":    fiber.StatusBadRequest,
			"data":    []string{},
		})
	}

	if err := br.bannerService.CreateBanner(file); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusInternalServerError,
			"data":    []string{},
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.HttpResponse{
		Message: "success",
		Code:    fiber.StatusCreated,
		Data:    []string{},
	})
}

func (br *bannerRouter) DeleteBanner(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := br.bannerService.DeleteBanner(id); err != nil {
		if err == domain.ErrBannerNotFound {
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

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Message: "success",
		Code:    fiber.StatusOK,
		Data:    []string{},
	})
}

func (br *bannerRouter) GetAllBanners(c *fiber.Ctx) error {
	banners, err := br.bannerService.GetAllBanners()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusInternalServerError,
			"data":    []string{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Message: "success",
		Code:    fiber.StatusOK,
		Data:    banners,
	})
}

func (br *bannerRouter) GetBanner(c *fiber.Ctx) error {
	id := c.Params("id")
	banner, err := br.bannerService.GetBanner(id)
	if err != nil {
		if err == domain.ErrBannerNotFound {
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

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Message: "success",
		Code:    fiber.StatusOK,
		Data:    banner,
	})
}

package packages

import (
	"peluang-server/domain"
	"peluang-server/dto"
	"peluang-server/internal/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type packagesRoute struct {
	packageService domain.PackageService
}

func NewRoute(app *fiber.App, packageService domain.PackageService) {
	packagesRoute := &packagesRoute{
		packageService,
	}

	api := app.Group("/api/v1/merchant/:merchant_id/packages")
	{
		api.Get("", packagesRoute.GetAllPackages)
		api.Get("/:id", packagesRoute.GetPackage)

	}

	protected := app.Group("/api/v1/merchant/:merchant_id/packages", middleware.Authenticate())
	{
		protected.Post("", packagesRoute.CreatePackage)
		protected.Put("/:id", packagesRoute.UpdatePackage)
		protected.Delete("/:id", packagesRoute.DeletePackage)

		protected.Post("/:package_id/lists", packagesRoute.CreateList)
		protected.Post("/:package_id/aditionals", packagesRoute.CreateAditional)
		protected.Delete("/lists/:id", packagesRoute.DeleteList)
		protected.Delete("/aditionals/:id", packagesRoute.DeleteAditional)
	}
}

func (pr *packagesRoute) CreatePackage(c *fiber.Ctx) error {
	merchantID := c.Params("merchant_id")
	packageData := new(dto.PackageRequest)
	if err := c.BodyParser(packageData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing body",
			"code":    fiber.StatusBadRequest,
			"data":    []string{},
		})
	}

	if err := pr.packageService.Insert(packageData, merchantID); err != nil {
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

func (pr *packagesRoute) GetAllPackages(c *fiber.Ctx) error {
	packages, err := pr.packageService.FindAll()
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
		Data:    packages,
	})
}

func (pr *packagesRoute) GetPackage(c *fiber.Ctx) error {
	id := c.Params("id")
	packageData, err := pr.packageService.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
			"code":    fiber.StatusNotFound,
			"data":    []string{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Message: "success",
		Code:    fiber.StatusOK,
		Data:    packageData,
	})
}

func (pr *packagesRoute) UpdatePackage(c *fiber.Ctx) error {
	id := c.Params("id")
	packageData := new(dto.PackageRequest)
	if err := c.BodyParser(packageData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing body",
			"code":    fiber.StatusBadRequest,
			"data":    []string{},
		})
	}

	if err := pr.packageService.Update(id, packageData); err != nil {
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

func (pr *packagesRoute) DeletePackage(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := pr.packageService.Delete(id); err != nil {
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

func (pr *packagesRoute) CreateList(c *fiber.Ctx) error {
	packageID := c.Params("package_id")
	listData := new(dto.CreateListRequest)
	if err := c.BodyParser(listData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing body",
			"code":    fiber.StatusBadRequest,
			"data":    []string{},
		})
	}

	if err := pr.packageService.InsertLists(listData.List, packageID); err != nil {
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

func (pr *packagesRoute) CreateAditional(c *fiber.Ctx) error {
	packageID := c.Params("package_id")
	aditionalData := new(dto.CreateAditionalRequest)
	if err := c.BodyParser(aditionalData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing body",
			"code":    fiber.StatusBadRequest,
			"data":    []string{},
		})
	}

	if err := pr.packageService.InsertAditionals(aditionalData.Aditional, packageID); err != nil {
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

func (pr *packagesRoute) DeleteList(c *fiber.Ctx) error {
	id := c.Params("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
			"code":    fiber.StatusBadRequest,
			"data":    []string{},
		})
	}

	if err := pr.packageService.DeleteList(uint(idUint)); err != nil {
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

func (pr *packagesRoute) DeleteAditional(c *fiber.Ctx) error {
	id := c.Params("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
			"code":    fiber.StatusBadRequest,
			"data":    []string{},
		})
	}

	if err := pr.packageService.DeleteAditional(uint(idUint)); err != nil {
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

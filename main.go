package main

import (
	"log"
	"peluang-server/internal/component"
	"peluang-server/internal/config"
	"peluang-server/internal/modules/banner"
	"peluang-server/internal/modules/merchant"
	"peluang-server/internal/modules/otp"
	"peluang-server/internal/modules/packages"
	"peluang-server/internal/modules/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	conf := config.NewConfig()

	db := component.GetDatabaseConnection(conf)

	component.Migrate(db)

	userRepository := user.NewRepository(db)
	userOtpRepository := otp.NewRepository(db)
	merchantRepository := merchant.NewRepository(db)
	bannerRepository := banner.NewRepository(db)
	packagesRepository := packages.NewRepository(db)

	userService := user.NewService(userRepository, userOtpRepository)
	merchantService := merchant.NewService(merchantRepository)
	bannerService := banner.NewService(bannerRepository)
	packagesService := packages.NewService(packagesRepository)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	user.NewRoute(app, userService)
	merchant.NewRoute(app, merchantService)
	banner.NewRoute(app, bannerService)
	packages.NewRoute(app, packagesService)

	log.Fatal(app.Listen(":" + conf.Srv.Port))
}

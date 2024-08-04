package router

import (
	travel "github.com/travel-api-main/api/app/travel"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/v1/", logger.New())

	travel.AuthRoutes(api)

}

package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	apps "github.com/travel-api-main/api/app"
	"github.com/travel-api-main/pkg/config"
)

func NewFiberServer() {
	app := fiber.New()

	app.Static("/public", "./public")
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,X-Request-Id,X-System-Auth,X-User-Auth",
		AllowOrigins:     "http://localhost:5173, https://travel.unreadhub.com/", // Specify allowed origins
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Fiber")
	})

	apps.SetupRoutes(app)

	if err := app.Listen(config.GetString("port")); err != nil {
		fmt.Println(err)
	}
}

package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/travel-api-main/api/app/travel/client"
	"github.com/travel-api-main/api/app/travel/handler"
	"github.com/travel-api-main/api/app/travel/services"
	"github.com/travel-api-main/api/app/travel/validation"
)

func AuthRoutes(r fiber.Router) {
	client := client.InitAuthClient()

	countryService := services.NewCountryService(client)
	travelService := services.NewTravelService(client)

	handler.InitCountry(countryService)
	handler.InitTravel(travelService)

	routes_country := r.Group("/country")
	routes_country.Get("/", handler.GetsCountry)
	routes_country.Get("/:id", handler.GetCountry)
	routes_country.Post("/", validation.ValidateCreateCountry, handler.CreateCountry)
	routes_country.Put("/:id", validation.ValidateUpdateCountry, handler.UpdateCountry)
	routes_country.Delete("/:id", handler.DeleteCountry)

	routes_travel := r.Group("/travelled_country")
	routes_travel.Get("/", handler.GetsTravel)
	routes_travel.Get("/:id", handler.GetTravel)
	routes_travel.Post("/", handler.CreateTravel)
	routes_travel.Put("/:id", handler.UpdateTravel)
	routes_travel.Delete("/:id", handler.DeleteTravel)

}

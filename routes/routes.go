package routes

import (
	"star-citizen/handlers"
	"star-citizen/repositories"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	repo := repositories.NewInMemoryPlanetRepository()

	planetHandler := handlers.NewPlanetHandler(repo)

	app.Post("/planets", planetHandler.AddPlanet)
	app.Get("/planets/:id?", planetHandler.GetPlanets)
	app.Put("/planets/:id", planetHandler.UpdatePlanet)
}

package handlers

import (
	"star-citizen/models"
	"star-citizen/repositories"
	"star-citizen/utils"

	"github.com/gofiber/fiber/v2"
)

type PlanetHandler struct {
	repo repositories.PlanetRepository
}

func NewPlanetHandler(repo repositories.PlanetRepository) *PlanetHandler {
	return &PlanetHandler{
		repo: repo,
	}
}

func (h *PlanetHandler) AddPlanet(c *fiber.Ctx) error {

	var planet models.Planet

	if err := c.BodyParser(&planet); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "result": "Cannot parse JSON"})
	}

	if err := utils.ValidatePlanet(planet); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "result": err.Error()})
	}

	hero, err := h.repo.AddPlanet(planet)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failure", "result": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "result": hero})
}
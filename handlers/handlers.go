package handlers

import (
	"fmt"
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

func (h *PlanetHandler) GetPlanets(c *fiber.Ctx) error {

	id := c.Params("id")

	if id == "" {

		planets, err := h.repo.ListPlanets()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failure", "result": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "result": planets})
	}

	planet, err := h.repo.GetPlanet(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failure", "result": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "result": planet})
}

func (h *PlanetHandler) UpdatePlanet(c *fiber.Ctx) error {

	// TODO: implement partial updates

	id := c.Params("id")

	var planet models.Planet

	if err := c.BodyParser(&planet); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "result": "Cannot parse JSON"})
	}

	if err := utils.ValidatePlanet(planet); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "result": err.Error()})
	}

	updated, err := h.repo.UpdatePlanet(id, planet)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failure", "result": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "result": updated})
}

func (h *PlanetHandler) DeletePlanet(c *fiber.Ctx) error {

	id := c.Params("id")

	err := h.repo.DeletePlanet(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failure", "result": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "result": fmt.Sprintf("planet with id %v deleted", id)})
}

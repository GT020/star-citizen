package handlers

import (
	"fmt"
	"star-citizen/models"
	"star-citizen/repositories"
	"star-citizen/utils"
	"strconv"

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

		minMass, _ := strconv.ParseFloat(c.Query("min_mass"), 64)
		maxMass, _ := strconv.ParseFloat(c.Query("max_mass"), 64)
		minRadius, _ := strconv.ParseFloat(c.Query("min_radius"), 64)
		maxRadius, _ := strconv.ParseFloat(c.Query("max_radius"), 64)

		filter := models.PlanetFilter{
			MinMass:   minMass,
			MaxMass:   maxMass,
			MinRadius: minRadius,
			MaxRadius: maxRadius,
		}

		planets, err := h.repo.ListPlanets(filter)

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

func (h *PlanetHandler) GetFuelEstimate(c *fiber.Ctx) error {

	id := c.Params("id")

	planet, err := h.repo.GetPlanet(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failure", "result": err.Error()})
	}

	crewCapacity, err := strconv.Atoi(c.Query("crew"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "result": "Cannot parse crew size"})
	}

	if crewCapacity <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failure", "result": "Crew size must be greater than 0"})
	}

	estimate := utils.CalculateFuel(planet, crewCapacity)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "result": estimate})
}

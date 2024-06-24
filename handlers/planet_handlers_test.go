package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"star-citizen/models"
	"star-citizen/repositories"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupTestApp() *fiber.App {
	app := fiber.New()
	repo := repositories.NewInMemoryPlanetRepository()

	planetHandler := NewPlanetHandler(repo)

	app.Post("/planets", planetHandler.AddPlanet)
	app.Get("/planets/:id?", planetHandler.GetPlanets)
	app.Put("/planets/:id", planetHandler.UpdatePlanet)
	app.Delete("/planets/:id", planetHandler.DeletePlanet)
	app.Get("/planets/:id/fuel", planetHandler.GetFuelEstimate)
	return app
}

func TestAddPlanet(t *testing.T) {
	app := setupTestApp()

	planet := models.Planet{
		Name:        "Earth",
		Description: "planet",
		Distance:    11,
		Radius:      1.0,
		Type:        models.Terrestrial,
		Mass:        1.0,
	}

	body, _ := json.Marshal(planet)
	req := httptest.NewRequest(http.MethodPost, "/planets", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	t.Logf("Parse %v", result)

	assert.Equal(t, "success", result["status"])
}

func TestGetPlanets(t *testing.T) {
	app := setupTestApp()

	planet := models.Planet{
		Name:        "Mars",
		Description: "Red",
		Distance:    20,
		Radius:      0.5,
		Type:        models.Terrestrial,
		Mass:        0.6,
	}
	body, _ := json.Marshal(planet)
	req := httptest.NewRequest(http.MethodPost, "/planets", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	app.Test(req)

	req = httptest.NewRequest(http.MethodGet, "/planets", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Equal(t, "success", result["status"])
}

func TestGetFuelEstimate(t *testing.T) {
	app := setupTestApp()

	planet := models.Planet{
		Name:        "Jupiter",
		Description: "Gaseous",
		Distance:    50,
		Radius:      1.0,
		Type:        models.GasGiant,
	}
	body, _ := json.Marshal(planet)
	req := httptest.NewRequest(http.MethodPost, "/planets", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	app.Test(req)

	req = httptest.NewRequest(http.MethodGet, "/planets/1/fuel?crew=10", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Equal(t, "success", result["status"])
}

func TestDeletePlanet(t *testing.T) {
	app := setupTestApp()

	planet := models.Planet{
		Name:        "Neptune",
		Description: "Gas",
		Distance:    30,
		Radius:      2.0,
		Type:        models.GasGiant,
	}
	body, _ := json.Marshal(planet)
	req := httptest.NewRequest(http.MethodPost, "/planets", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	app.Test(req)

	req = httptest.NewRequest(http.MethodDelete, "/planets/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Equal(t, "success", result["status"])
}

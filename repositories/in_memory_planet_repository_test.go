package repositories

import (
	"star-citizen/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddPlanet(t *testing.T) {
	repo := NewInMemoryPlanetRepository()

	planet := models.Planet{
		Name:        "Earth",
		Description: "Blue",
		Distance:    1,
		Radius:      1.0,
		Type:        models.Terrestrial,
		Mass:        1.0,
	}

	created, err := repo.AddPlanet(planet)

	assert.NoError(t, err)
	assert.Equal(t, "1", created.ID)
}

func TestListPlanets(t *testing.T) {
	repo := NewInMemoryPlanetRepository()

	planet1 := models.Planet{
		Name:        "Earth",
		Description: "Home",
		Distance:    1,
		Radius:      1.0,
		Type:        models.Terrestrial,
		Mass:        1.0,
	}

	planet2 := models.Planet{
		Name:        "Mars",
		Description: "Red",
		Distance:    2,
		Radius:      0.5,
		Type:        models.Terrestrial,
		Mass:        0.6,
	}

	repo.AddPlanet(planet1)
	repo.AddPlanet(planet2)

	filter := models.PlanetFilter{
		MinMass: 0.5,
		MaxMass: 1.0,
	}

	planets, err := repo.ListPlanets(filter)

	assert.NoError(t, err)
	assert.Len(t, planets, 2)
}

func TestGetPlanet(t *testing.T) {
	repo := NewInMemoryPlanetRepository()

	planet := models.Planet{
		Name:        "Earth",
		Description: "Green",
		Distance:    1,
		Radius:      1.0,
		Type:        models.Terrestrial,
		Mass:        1.0,
	}

	repo.AddPlanet(planet)

	retrieved, err := repo.GetPlanet("1")

	assert.NoError(t, err)
	assert.Equal(t, "Earth", retrieved.Name)
}

func TestUpdatePlanet(t *testing.T) {
	repo := NewInMemoryPlanetRepository()

	planet := models.Planet{
		Name:        "Earth",
		Description: "Home",
		Distance:    1,
		Radius:      1.0,
		Type:        models.Terrestrial,
		Mass:        1.0,
	}

	repo.AddPlanet(planet)

	planet.Name = "Updated Earth"
	updated, err := repo.UpdatePlanet("1", planet)

	assert.NoError(t, err)
	assert.Equal(t, "Updated Earth", updated.Name)
}

func TestDeletePlanet(t *testing.T) {
	repo := NewInMemoryPlanetRepository()

	planet := models.Planet{
		Name:        "Earth",
		Description: "Our home planet",
		Distance:    1,
		Radius:      1.0,
		Type:        models.Terrestrial,
		Mass:        1.0,
	}

	repo.AddPlanet(planet)

	err := repo.DeletePlanet("1")

	assert.NoError(t, err)
}

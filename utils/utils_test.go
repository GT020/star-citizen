package utils

import (
	"star-citizen/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateFuel(t *testing.T) {
	planet := models.Planet{
		Distance: 100,
		Radius:   1.0,
		Type:     models.Terrestrial,
		Mass:     1.0,
	}

	fuel := CalculateFuel(planet, 10)
	assert.Equal(t, 1000.0, fuel)
}

func TestValidatePlanet(t *testing.T) {
	planet := models.Planet{
		Name:        "Mars",
		Description: "Blue",
		Distance:    11,
		Radius:      1.0,
		Type:        models.Terrestrial,
		Mass:        1.0,
	}

	err := ValidatePlanet(planet)
	assert.NoError(t, err)

	planet.Type = "invalid"
	err = ValidatePlanet(planet)
	assert.Error(t, err)
}

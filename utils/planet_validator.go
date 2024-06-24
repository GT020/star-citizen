package utils

import (
	"errors"
	"star-citizen/models"
)

func ValidatePlanet(planet models.Planet) error {
	if planet.Name == "" || planet.Description == "" {
		return errors.New("invalid planet data")
	}

	if planet.Distance <= 10 || planet.Distance >= 1000 {
		return errors.New("invalid distance for planet")
	}

	if planet.Radius <= 0.1 || planet.Radius >= 10 {
		return errors.New("invalid radius for planet")
	}

	if err := planet.Type.IsValid(); err != nil {
		return err
	}

	if planet.Type == models.Terrestrial && (planet.Mass <= 0.1 || planet.Mass >= 10) {
		return errors.New("invalid mass for Terrestrial planet")
	}

	if planet.Type == models.GasGiant && planet.Mass != 0 {
		return errors.New("mass will be empty for gas giant")
	}

	return nil
}

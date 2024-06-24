package utils

import "star-citizen/models"

func CalculateFuel(planet models.Planet, crewCapacity int) float64 {
	d := float64(planet.Distance)
	g := planet.GetGravity()

	return d / (g * g) * float64(crewCapacity)
}

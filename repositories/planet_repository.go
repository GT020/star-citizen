package repositories

import (
	"star-citizen/models"
)

type PlanetRepository interface {
	AddPlanet(Planet models.Planet) (models.Planet, error)
	ListPlanets() ([]models.Planet, error)
	GetPlanet(id string) (models.Planet, error)
	UpdatePlanet(id string, Planet models.Planet) (models.Planet, error)
	DeletePlanet(id string) error
}

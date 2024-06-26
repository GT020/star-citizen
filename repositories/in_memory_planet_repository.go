package repositories

import (
	"fmt"
	"star-citizen/models"
)

type InMemoryPlanetRepository struct {
	planets   map[string]models.Planet
	currentID int
}

func NewInMemoryPlanetRepository() *InMemoryPlanetRepository {
	return &InMemoryPlanetRepository{
		planets:   make(map[string]models.Planet),
		currentID: 0,
	}
}

func (r *InMemoryPlanetRepository) AddPlanet(Planet models.Planet) (models.Planet, error) {

	if _, ok := r.planets[Planet.ID]; ok {
		return models.Planet{}, fmt.Errorf("planet with id %v already exists", Planet.ID)
	}

	r.currentID++

	Planet.ID = fmt.Sprintf("%v", r.currentID)

	r.planets[Planet.ID] = Planet

	return Planet, nil
}
func (r *InMemoryPlanetRepository) ListPlanets(filter models.PlanetFilter) ([]models.Planet, error) {

	res := make([]models.Planet, 0, len(r.planets))

	for _, planet := range r.planets {
		if (filter.MinMass == 0 || planet.Mass >= filter.MinMass) &&
			(filter.MaxMass == 0 || planet.Mass <= filter.MaxMass) &&
			(filter.MinRadius == 0 || planet.Radius >= filter.MinRadius) &&
			(filter.MaxRadius == 0 || planet.Radius <= filter.MaxRadius) {
			res = append(res, planet)
		}
	}

	return res, nil
}

func (r *InMemoryPlanetRepository) GetPlanet(id string) (models.Planet, error) {

	planet, ok := r.planets[id]

	if !ok {
		return models.Planet{}, fmt.Errorf("planet with id %v does not exists", id)
	}

	return planet, nil

}

func (r *InMemoryPlanetRepository) UpdatePlanet(id string, Planet models.Planet) (models.Planet, error) {
	_, ok := r.planets[id]

	if !ok {
		return models.Planet{}, fmt.Errorf("planet with id %v does not exists", id)
	}

	Planet.ID = id
	r.planets[id] = Planet

	return Planet, nil
}

func (r *InMemoryPlanetRepository) DeletePlanet(id string) error {
	_, ok := r.planets[id]

	if !ok {
		return fmt.Errorf("planet with id %v does not exists", id)
	}

	delete(r.planets, id)

	return nil
}

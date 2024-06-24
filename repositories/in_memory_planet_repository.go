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
		currentID: 1,
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
func (r *InMemoryPlanetRepository) ListPlanets() ([]models.Planet, error) {

	res := make([]models.Planet, 0, len(r.planets))

	count := 0

	for _, planet := range r.planets {
		res[count] = planet
		count++
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

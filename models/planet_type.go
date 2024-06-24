package models

import "errors"

type PlanetType string

const (
	GasGiant    PlanetType = "gas-giant"
	Terrestrial PlanetType = "terrestrial"
)

func (pt PlanetType) IsValid() error {
	switch pt {
	case GasGiant, Terrestrial:
		return nil
	}
	return errors.New("invalid planet type")
}

package models

type Planet struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Distance    int        `json:"distance"`
	Radius      float64    `json:"radius"`
	Mass        float64    `json:"mass,omitempty"`
	Type        PlanetType `json:"type"`
}

func (p *Planet) GetGravity() float64 {

	switch p.Type {
	case GasGiant:
		return (0.5 / (p.Radius * p.Radius))

	case Terrestrial:
		return (p.Mass / (p.Radius * p.Radius))

	default:
		return 0.0
	}
}

package environment

import "github.com/jecolasurdo/marsrover/pkg/coordinate"

// Plateau is a rectangular maritan environment.
type Plateau struct {
	dimensions coordinate.Point
}

// NewPlateau instantiates a new Plateau and returns a reference to that
// instance.
func NewPlateau(dimensions coordinate.Point) *Plateau {
	return &Plateau{
		dimensions: dimensions,
	}
}

// GetDimensions returns the dimension of the environment.
func (p *Plateau) GetDimensions() coordinate.Point {
	return p.dimensions
}

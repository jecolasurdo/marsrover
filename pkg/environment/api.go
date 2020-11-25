package environment

import "github.com/jecolasurdo/marsrover/pkg/coordinate"

// Plateau is a rectangular maritan environment.
type Plateau struct {
	size coordinate.Point
}

// NewPlateau instantiates a new Plateau and returns a reference to that
// instance.
func NewPlateau(size coordinate.Point) *Plateau {
	return &Plateau{
		size: size,
	}
}

// GetDimensions returns the dimension of the environment.
func (p *Plateau) GetDimensions() coordinate.Point {
	panic("not implemented")
}

package environment

import (
	"github.com/jecolasurdo/marsrover/pkg/coordinate"
	"github.com/jecolasurdo/marsrover/pkg/object/objectiface"
)

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

// PlaceObject inserts a new object into the environment at some position.
// The environment will enforce unique object ID's for consistency.
func (p *Plateau) PlaceObject(object objectiface.Objecter, position coordinate.Point) error {
	panic("not implemented")
}

// MoveObject moves an object from one point in the environment to another.
func (p *Plateau) MoveObject(object objectiface.Objecter, newPosition coordinate.Point) error {
	panic("not implemented")
}

// ShowObjects returns a sparse map of points within the terrain that
// contain objects.
func ShowObjects() map[coordinate.Point][]objectiface.Objecter {
	panic("not implemented")
}

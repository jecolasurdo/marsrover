package environment

import (
	"github.com/jecolasurdo/marsrover/pkg/coordinate"
	"github.com/jecolasurdo/marsrover/pkg/object/objectiface"
)

type objectStore map[coordinate.Point][]objectiface.Objecter

// Plateau is a rectangular maritan environment.
type Plateau struct {
	dimensions coordinate.Point
	objects    objectStore
}

// NewPlateau instantiates a new Plateau and returns a reference to that
// instance.
func NewPlateau(dimensions coordinate.Point) *Plateau {
	return &Plateau{
		dimensions: dimensions,
		objects:    make(objectStore),
	}
}

// GetDimensions returns the dimension of the environment.
func (p *Plateau) GetDimensions() coordinate.Point {
	return p.dimensions
}

// PlaceObject inserts a new object into the environment at some position.
// The environment will enforce unique object ID's for consistency.
func (p *Plateau) PlaceObject(object objectiface.Objecter, position coordinate.Point) error {
	if objectList, found := p.objects[position]; found {
		objectList = append(objectList, object)
	} else {
		p.objects[position] = []objectiface.Objecter{object}
	}
	return nil
}

// MoveObject moves an object from one point in the environment to another.
func (p *Plateau) MoveObject(object objectiface.Objecter, newPosition coordinate.Point) error {
	panic("not implemented")
}

// ShowObjects returns a sparse map of points within the terrain that
// contain objects.
func (p *Plateau) ShowObjects() map[coordinate.Point][]objectiface.Objecter {
	return p.objects
}

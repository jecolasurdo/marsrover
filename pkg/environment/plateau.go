package environment

import (
	"fmt"

	"github.com/jecolasurdo/marsrover/pkg/coordinate"
	"github.com/jecolasurdo/marsrover/pkg/object/objectiface"
)

type objectStore map[coordinate.Point][]objectiface.Objecter

// Plateau is a rectangular martian environment.
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
	if object == nil {
		return fmt.Errorf("a nil object cannot be placed in the environment")
	}

	if position.X > p.dimensions.X || position.Y > p.dimensions.Y ||
		position.Y < 0 || position.X < 0 {
		return fmt.Errorf("an object cannot be placed outside the bounds of the environment")
	}

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

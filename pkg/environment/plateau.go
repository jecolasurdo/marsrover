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

// PlaceObject attempts to insert a new object into the environment at the
// specified position.
//
// The environment enforces the following rules when placing objects:
//   - nil objects cannot be placed in the environment.
//   - Objects can only be placed within the bounds of the environment.
//   - Each object ID can only exist once within the environment.
//   - Multiple objects are allowed to share the same position in the
//     environment. (e.g. It is each object's responsibility to determine
//     whether or not it can occupy the same space as another object; this is
//     not enforced by the environment itself.)
//
// In each of the above cases, PlaceObject will return an error if the rule is
// violated.
//
// PlaceObject will return nil if the object has been successfully placed at the
// specified position.
func (p *Plateau) PlaceObject(object objectiface.Objecter, position coordinate.Point) error {
	if object == nil {
		return fmt.Errorf("a nil object cannot be placed in the environment")
	}

	err := p.verifyPositionIsLegal(position)
	if err != nil {
		return err
	}

	for _, existingObjects := range p.objects {
		for _, existingObject := range existingObjects {
			if existingObject.ID() == object.ID() {
				return fmt.Errorf("object with ID '%s' already exists within the environment", object.ID())
			}
		}
	}

	if objectList, found := p.objects[position]; found {
		objectList = append(objectList, object)
	} else {
		p.objects[position] = []objectiface.Objecter{object}
	}
	return nil
}

// RecordMovement records the movement an object from one position in the
// environment to another.
func (p *Plateau) RecordMovement(object objectiface.Objecter, newPosition coordinate.Point) error {
	if object == nil {
		return fmt.Errorf("cannot record the movement of a nil object")
	}

	err := p.verifyPositionIsLegal(newPosition)
	if err != nil {
		return err
	}
	return nil
}

// ShowObjects returns a sparse map of points within the terrain that
// contain objects.
func (p *Plateau) ShowObjects() map[coordinate.Point][]objectiface.Objecter {
	return p.objects
}

func (p *Plateau) verifyPositionIsLegal(position coordinate.Point) error {
	if position.X > p.dimensions.X || position.Y > p.dimensions.Y ||
		position.Y < 0 || position.X < 0 {
		return fmt.Errorf("an object cannot be placed outside the bounds of the environment")
	}
	return nil
}

package environment

import (
	"github.com/jecolasurdo/marsrover/pkg/environment/environmentiface"
	"github.com/jecolasurdo/marsrover/pkg/environment/environmenttypes"
	"github.com/jecolasurdo/marsrover/pkg/objects/objectiface"
	"github.com/jecolasurdo/marsrover/pkg/spatial"
)

type objectStore map[spatial.Point][]objectiface.Objecter

// Plateau is a curiously rectangular martian environment.
type Plateau struct {
	dimensions spatial.Point
	objects    objectStore
}

// NewPlateau instantiates a new Plateau and returns a reference to that
// instance.
func (Plateau) NewPlateau(dimensions spatial.Point) *Plateau {
	return &Plateau{
		dimensions: dimensions,
		objects:    make(objectStore),
	}
}

// GetDimensions returns the dimension of the environment.
func (p *Plateau) GetDimensions() spatial.Point {
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
func (p *Plateau) PlaceObject(object objectiface.Objecter, position spatial.Point) error {
	if object == nil {
		return ErrNilObject()
	}

	err := p.verifyPositionIsLegal(position)
	if err != nil {
		return err
	}

	if found, _ := p.FindObject(object); found {
		return ErrObjectAlreadyExists(object)
	}

	p.placeObjectUnchecked(object, position)
	return nil
}

// RecordMovement records the movement of an object from one position in the
// environment to another.
func (p *Plateau) RecordMovement(object objectiface.Objecter, newPosition spatial.Point) error {
	if object == nil {
		return ErrNilObject()
	}

	err := p.verifyPositionIsLegal(newPosition)
	if err != nil {
		return err
	}

	found, objectPosition := p.FindObject(object)
	if !found {
		return ErrObjectDoesNotExist(object)
	}

	p.removeObjectUnchecked(objectPosition.Object)
	p.placeObjectUnchecked(object, newPosition)

	return nil
}

// ShowObjects returns a sparse map of points within the terrain that
// contain objects.
func (p *Plateau) ShowObjects() map[spatial.Point][]objectiface.Objecter {
	return p.objects
}

// FindObject searches the environment for an object (via the object's ID)
// and if the object is found, returns true and the object and its position.
// If the object is not found in the environment, FindObject returns false.
func (p *Plateau) FindObject(objectToFind objectiface.Objecter) (bool, *environmenttypes.ObjectPosition) {
	for position, objects := range p.objects {
		for _, object := range objects {
			if object.ID() == objectToFind.ID() {
				return true, &environmenttypes.ObjectPosition{
					Position: position,
					Object:   object,
				}
			}
		}
	}
	return false, nil
}

// InspectPosition attempts to return any objects that may be present at
// a specified position.
//
// If there are no objects present at the specified position, this method
// will return false, either an empty or nil slice, and a nil error.
//
// If there are objects present at the specified position, this method will
// return true, a non-empty list of objects, and a nil error.
//
// If the requested position does not exist within the plateau, this method will
// return false, nil, and an error.
func (p *Plateau) InspectPosition(positionToInspect spatial.Point) (bool, []objectiface.Objecter, error) {
	err := p.verifyPositionIsLegal(positionToInspect)
	if err != nil {
		return false, nil, err
	}

	for knownPosition, objects := range p.objects {
		if knownPosition == positionToInspect {
			if objects != nil && len(objects) > 0 {
				return true, objects, nil
			}
			break
		}
	}
	return false, nil, nil
}

func (p *Plateau) verifyPositionIsLegal(position spatial.Point) error {
	if position.X > p.dimensions.X || position.Y > p.dimensions.Y ||
		position.Y < 0 || position.X < 0 {
		return ErrPositionOutsideBounds(position)
	}
	return nil
}

// removeObjectUnchecked removes an object from the environment without checking
// if the object already exists nor performing any other validity checks.
func (p *Plateau) removeObjectUnchecked(object objectiface.Objecter) {
	for position, existingObjects := range p.objects {
		objectsAtCurrentPosition := []objectiface.Objecter{}
		for _, existingObject := range existingObjects {
			if existingObject.ID() != object.ID() {
				objectsAtCurrentPosition = append(objectsAtCurrentPosition, existingObject)
			}
		}
		if len(objectsAtCurrentPosition) == 0 {
			delete(p.objects, position)
		} else {
			p.objects[position] = objectsAtCurrentPosition
		}
	}
}

// placeObjectUnchecked places the specified object at the specified position.
// This method places an object into the environment without checking the validity
// of the specified coordinates, without checking if the object is nil, and
// without checking if the object already exists elsewhere in the environment.
func (p *Plateau) placeObjectUnchecked(object objectiface.Objecter, newPosition spatial.Point) {
	if objectList, found := p.objects[newPosition]; found {
		objectList = append(objectList, object)
		p.objects[newPosition] = objectList
	} else {
		p.objects[newPosition] = []objectiface.Objecter{object}
	}
}

// enforce that Plateau implements Environmenter
var _ environmentiface.Environmenter = (*Plateau)(nil)

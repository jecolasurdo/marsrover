package environmentiface

import (
	"github.com/jecolasurdo/marsrover/pkg/coordinate"
	"github.com/jecolasurdo/marsrover/pkg/object/objectiface"
)

// Environmenter is anything that can describe an environment.
type Environmenter interface {
	// GetDimenstions returns the dimension of the environment.
	GetDimensions() coordinate.Point

	// Place object inserts a new object into the environment at some position.
	// The environment will enforce unique object ID's for consistency.
	PlaceObject(objectiface.Objecter, coordinate.Point) error

	// MoveObject moves an object from one point in the environment to another.
	MoveObject(objectiface.Objecter, coordinate.Point) error

	// ShowObjects returns a sparse map of points within the terrain that
	// contain objects.
	ShowObjects() map[coordinate.Point][]objectiface.Objecter
}

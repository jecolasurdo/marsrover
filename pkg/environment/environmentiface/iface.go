package environmentiface

import (
	"github.com/jecolasurdo/marsrover/pkg/coordinate"
	"github.com/jecolasurdo/marsrover/pkg/environment/environmenttypes"
	"github.com/jecolasurdo/marsrover/pkg/object/objectiface"
)

// Environmenter is anything that can describe an environment.
type Environmenter interface {
	// GetDimenstions returns the dimension of the environment.
	GetDimensions() coordinate.Point

	// PlaceObject inserts a new object into the environment at some position.
	// The environment will enforce unique object ID's for consistency.
	PlaceObject(objectiface.Objecter, coordinate.Point) error

	// RecordMovement records the movement an object from one position in the
	// environment to another.
	RecordMovement(objectiface.Objecter, coordinate.Point) error

	// ShowObjects returns a sparse map of points within the terrain that
	// contain objects.
	ShowObjects() map[coordinate.Point][]objectiface.Objecter

	// FindObject searches the environment for an object (via the object's ID)
	// and if the object is found, returns true and the object and its position.
	// If the object is not found in the environment, FindObject returns false.
	FindObject(objectiface.Objecter) (bool, *environmenttypes.ObjectPosition)
}

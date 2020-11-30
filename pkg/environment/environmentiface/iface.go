// Package environmentiface provides contracts for basic environment behavior.
package environmentiface

import (
	"github.com/jecolasurdo/marsrover/pkg/environment/environmenttypes"
	"github.com/jecolasurdo/marsrover/pkg/objects/objectiface"
	"github.com/jecolasurdo/marsrover/pkg/spatial"
)

// EnvironmentBuilder is anything that knows how to construct an abstract
// environment.
type EnvironmentBuilder interface {
	NewEnvironment(spatial.Point) (Environmenter, error)
}

// Environmenter is anything that can describe an environment.
type Environmenter interface {
	// GetDimenstions returns the dimension of the environment.
	GetDimensions() spatial.Point

	// PlaceObject inserts a new object into the environment at some position.
	// The environment will enforce unique object ID's for consistency.
	PlaceObject(objectiface.Objecter, spatial.Point) error

	// RecordMovement records the movement an object from one position in the
	// environment to another.
	RecordMovement(objectiface.Objecter, spatial.Point) error

	// ShowObjects returns a sparse map of points within the terrain that
	// contain objects.
	ShowObjects() map[spatial.Point][]objectiface.Objecter

	// FindObject searches the environment for an object (via the object's ID)
	// and if the object is found, returns true and the object and its position.
	// If the object is not found in the environment, FindObject returns false.
	FindObject(objectiface.Objecter) (bool, *environmenttypes.ObjectPosition)

	// InspectPosition attempts to return any objects that may be present at
	// a specified position.
	//
	// If there are no objects present at the specified position, this method
	// must return false, either an empty or nil slice, and a nil error.
	//
	// If there are objects present at the specified position, this method must
	// return true, a non-empty list of objects, and a nil error.
	//
	// If the requested position does not exist within the defined environment,
	// the method must return false, nil, and an error.
	InspectPosition(spatial.Point) (bool, []objectiface.Objecter, error)
}

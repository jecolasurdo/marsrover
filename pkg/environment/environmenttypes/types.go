// Package environmenttypes provides public types associated with environment behavior.
package environmenttypes

import (
	"github.com/jecolasurdo/marsrover/pkg/objects/objectiface"
	"github.com/jecolasurdo/marsrover/pkg/spatial"
)

// ObjectPosition represents the position of an object.
type ObjectPosition struct {
	Object   objectiface.Objecter
	Position spatial.Point
}

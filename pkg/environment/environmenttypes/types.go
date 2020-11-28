package environmenttypes

import (
	"github.com/jecolasurdo/marsrover/pkg/coordinate"
	"github.com/jecolasurdo/marsrover/pkg/object/objectiface"
)

// ObjectPosition represents the position of an object.
type ObjectPosition struct {
	Object   objectiface.Objecter
	Position coordinate.Point
}

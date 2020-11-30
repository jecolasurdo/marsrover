package environment

import (
	"fmt"

	"github.com/jecolasurdo/marsrover/pkg/objects/objectiface"
	"github.com/jecolasurdo/marsrover/pkg/spatial"
)

// ErrNilObject occurs if a nil object is used illegally.
func ErrNilObject() error {
	return fmt.Errorf("the environment cannot interact with a nil object")
}

// ErrObjectAlreadyExists occurs if an object already exists where/when it
// shouldn't.
func ErrObjectAlreadyExists(object objectiface.Objecter) error {
	return fmt.Errorf("object with ID '%s' already exists within the environment", object.ID())
}

// ErrObjectDoesNotExist occurs if an object does not exist where/when it should.
func ErrObjectDoesNotExist(object objectiface.Objecter) error {
	return fmt.Errorf("object with ID '%s' does not exist within the environment", object.ID())
}

// ErrPositionOutsideBounds occurs if a position is outside of the bounds of the
// environment in a situation where that is prohibited.
func ErrPositionOutsideBounds(position spatial.Point) error {
	return fmt.Errorf("position '%v' is outside the bounds of the environment", position)
}

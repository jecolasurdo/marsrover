package objects

import (
	"fmt"

	"github.com/jecolasurdo/marsrover/pkg/spatial"
)

// ErrRoverExpelledFromEnvironment occurs if the rover's underlaying environment
// no longer recognizes the rover as existing.
func ErrRoverExpelledFromEnvironment(rover *Rover) error {
	return fmt.Errorf("rover '%v' is no longer recognised by its environment", rover.ID())
}

// ErrRoverIncompatibleObjectDetected is returned if a rover detects an
// incompatible object at some position within its environment.
func ErrRoverIncompatibleObjectDetected(position spatial.Point) error {
	return fmt.Errorf("an incompatible object was dectected at position '%v'", position)
}

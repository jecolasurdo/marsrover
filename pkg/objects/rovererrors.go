package objects

import (
	"fmt"
)

// ErrRoverExpelledFromEnvironment occurs if the rover's underlaying environment
// no longer recognizes the rover as existing.
func ErrRoverExpelledFromEnvironment(rover *Rover) error {
	return fmt.Errorf("rover '%v' is no longer recognised by its environment", rover.ID())
}

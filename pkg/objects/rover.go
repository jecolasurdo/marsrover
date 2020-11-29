package objects

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jecolasurdo/marsrover/pkg/coordinate"
	"github.com/jecolasurdo/marsrover/pkg/environment/environmentiface"
)

// A Rover is a vehicle that traverses an environment.
type Rover struct {
	id      string
	env     environmentiface.Environmenter
	heading Heading
}

// LaunchRover initializes a new rover, and attempts to place it within the
// environment. The placement within the environment is subject to the spatial
// rules of the supplied environment. If the rover cannot be placed at the
// supplied position (due to the environment's rules), an error will be
// returned, and the rover will not initialize (it will be nil).
func LaunchRover(heading Heading, position coordinate.Point, environment environmentiface.Environmenter) (*Rover, error) {
	rover := &Rover{
		id:      uuid.New().String(),
		env:     environment,
		heading: heading,
	}

	err := environment.PlaceObject(rover, position)
	if err != nil {
		return nil, err
	}
	return rover, nil
}

// ID returns a string that uniquely identifies this Rover instance.
func (r *Rover) ID() string {
	return r.id
}

// CurrentPosition returns the rover's current position within its environment.
// An error can also be returned if the rover has somehow become removed from
// its environment (such a situation can occur depending on the rules of the
// environment in which the rover is currently operating).
func (r *Rover) CurrentPosition() (*coordinate.Point, error) {
	found, objectPosition := r.env.FindObject(r)
	if !found {
		return nil, fmt.Errorf("this rover no longer exists within its environment")
	}
	return &objectPosition.Position, nil
}

// CurrentHeading returns the rover's current heading.
func (r *Rover) CurrentHeading() Heading {
	return r.heading
}

// ChangeHeading updates the rover's current heading according to a specified
// direction (left of right).
func (r *Rover) ChangeHeading(direction Direction) {
	panic("not implemented")
}

// Move attempts to move the rover forward one unit in its current heading.
//
// If the move succeeds, the rover has successfully changed positions, and this
// method will return nil
//
// However, the move can fail in a few situations.
// - If the movement in the current heading would violate the rules of the
//   current environment. (typically this means going out of bounds, but the
//   behavior can change depending on the rules of a particular environment.
// - If the next position would result in moving to a space already occupied
//   by another object in the environment.
// If the move fails, an error will be returned, and the rover will retain its
// previous heading and position.
func (r *Rover) Move() error {
	panic("not implemented")
}

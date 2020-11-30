// Package roveriface provides contracts for basic rover behavior.
package roveriface

import (
	"github.com/jecolasurdo/marsrover/pkg/environment/environmentiface"
	"github.com/jecolasurdo/marsrover/pkg/objects/objectiface"
	"github.com/jecolasurdo/marsrover/pkg/spatial"
)

// RoverBuilder is anything that knows how to construc an abstract rover.
type RoverBuilder interface {
	// LaunchRover initializes a new rover, and attempts to place it within the
	// environment.
	LaunchRover(spatial.Heading, spatial.Point, environmentiface.Environmenter) (RoverAPI, error)
}

// RoverAPI represents anything that can behave like a rover.
type RoverAPI interface {
	objectiface.Objecter

	// CurrentPosition must report the rover's current position within its
	// environment.
	// The reported position may be nil if an error is encountered.
	CurrentPosition() (*spatial.Point, error)

	// CurrentHeading must report the rover's current heading.
	CurrentHeading() spatial.Heading

	// ChangeHeading must update the rover's current heading.
	ChangeHeading(spatial.Direction)

	// Move attempts to move the Rover according to its implementation specific
	// rules.
	Move() error
}

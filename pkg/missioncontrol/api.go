// Package missioncontrol provides capabilities for executing missons.
package missioncontrol

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jecolasurdo/marsrover/pkg/environment/environmentiface"
	"github.com/jecolasurdo/marsrover/pkg/objects/roveriface"
	"github.com/jecolasurdo/marsrover/pkg/spatial"
)

// A Mission is an interaction between objects (such as rovers) and an
// environment (such as a martian plateau).
type Mission struct {
	envBuilder   environmentiface.EnvironmentBuilder
	roverBuilder roveriface.RoverBuilder
}

// NewMission constructs a new mission.
func NewMission(envBuilder environmentiface.EnvironmentBuilder, roverBuilder roveriface.RoverBuilder) *Mission {
	return &Mission{
		envBuilder,
		roverBuilder,
	}
}

// ExecuteMission executes a mission in an environment according to the supplied
// commands.
//
// This method will immediately halt the mission and return an error if there
// is any problem detected within the mission.
func (m *Mission) ExecuteMission(commands []string) ([]string, error) {
	roverstats := []string{}
	env, commands, err := m.EstablishEnvironment(commands)
	if err != nil {
		return nil, err
	}

	for len(commands) > 0 {
		stats := ""
		stats, commands, err = m.DeployAndNavigateRover(env, commands)
		if err != nil {
			return nil, err
		}
		roverstats = append(roverstats, stats)
	}

	return roverstats, nil
}

// EstablishEnvironment attempts to construct a new environment based on the
// supplied commands.
//
// At least one command must be supplied to this method, and only the first
// command is observed. If successful, the method will consume the first command,
// and return the remaining unused commands for further processing by the
// caller.
//
// If the method succeeds, it will return an environment, and a list of all
// commands that it did not consume during its operation.
//
// If the method fails, only an error is returned.
func (m *Mission) EstablishEnvironment(commands []string) (environmentiface.Environmenter, []string, error) {
	if len(commands) == 0 {
		return nil, nil, ErrParsingRoverCommand("")
	}

	envCommand := commands[0]
	coords := strings.Split(envCommand, " ")
	if len(coords) != 2 {
		return nil, nil, ErrParsingEnvironmentCommand(envCommand)
	}

	x, err := strconv.Atoi(coords[0])
	if err != nil {
		return nil, nil, ErrParsingEnvironmentCommand(envCommand)
	}

	y, err := strconv.Atoi(coords[1])
	if err != nil {
		return nil, nil, ErrParsingEnvironmentCommand(envCommand)
	}

	env, err := m.envBuilder.NewEnvironment(spatial.NewPoint(x, y))
	if err != nil {
		return nil, nil, err
	}

	return env, commands[1:], nil
}

// DeployAndNavigateRover attempts to deploy a rover within an environment, and
// then navigate that rover according to supplied commands.
//
// At least two commands must be supplied to this method, and only the two
// commands are observed. If successful, the method will consume the first two
// commands and return the remaining unused commands for further processing by
// the caller. See PlaceRoverInEnvironment and NavigateRover for information
// about each of the commandds expected by this method.
//
// If the method succeeds, the status of the rover is returned along with a
// list of remaining commands. See NavigateRover for information about the
// format of the rover's status message.
//
// If the method fails, then the only an error is returned.
func (m *Mission) DeployAndNavigateRover(env environmentiface.Environmenter, commands []string) (string, []string, error) {
	if len(commands) < 2 {
		return "", nil, ErrParsingRoverCommand("expected at least two commands")
	}

	rover, commands, err := m.PlaceRoverInEnvironment(env, commands)
	if err != nil {
		return "", nil, err
	}

	return m.NavigateRover(rover, commands)
}

// PlaceRoverInEnvironment attempts to establish a new rover and place it
// within the specified environment.
//
// At least one command must be supplied to this method, and only the first
// command is observed. If successful, the method will consume the first command,
// and return the remaining unused commands for further processing by the
// caller.
//
// The command must be formatted as a space delimited string with
// three fields in the following order: 'x y h' where x is an x position, y
// is a y position, and h is a heading expressed as a cardinal value N, E, S,
// or W.
//
// If the method fails to place a rover in its environment, then only an error
// is returned.
func (m *Mission) PlaceRoverInEnvironment(env environmentiface.Environmenter, commands []string) (roveriface.RoverAPI, []string, error) {
	if len(commands) < 1 {
		return nil, nil, ErrParsingRoverCommand("")
	}

	positionCommands := strings.Split(commands[0], " ")
	if len(positionCommands) != 3 {
		return nil, nil, ErrParsingRoverCommand(commands[0])
	}

	x, err := strconv.Atoi(positionCommands[0])
	if err != nil {
		return nil, nil, ErrParsingEnvironmentCommand(commands[0])
	}

	y, err := strconv.Atoi(positionCommands[1])
	if err != nil {
		return nil, nil, ErrParsingEnvironmentCommand(commands[0])
	}

	heading := spatial.HeadingFromString(positionCommands[2])
	if heading == spatial.HeadingUnknown {
		return nil, nil, ErrParsingEnvironmentCommand(commands[0])
	}

	rover, err := m.roverBuilder.LaunchRover(heading, spatial.NewPoint(x, y), env)
	if err != nil {
		return nil, nil, err
	}

	return rover, commands[1:], nil
}

// NavigateRover attempts to maneaver a rover in an environment according to
// a supplied command.
//
// At least one command must be supplied to this method, and only the first
// command is observed. If successful, the method will consume the first command,
// and return the remaining unused commands for further processing by the
// caller.
//
// Valid values for the command are L, which represents a 90 degree turn to the
// left, R, which represents a 90 degree turn to the right, and M, which
// represents a move forward in the rover's current heading.
//
// If the method succeeds, then it returns the current status of the rover along
// with a list of remaining commands.
//
// The status of the rover is expressed as a single string with three values as
// follow: "{x coordinate} {y coordinate} {heading}"
//
// If the method fails when navigating the rover, then only an error is
// returned.
func (m *Mission) NavigateRover(rover roveriface.RoverAPI, commands []string) (string, []string, error) {
	for _, navigationCommand := range commands {
		if navigationCommand == "M" {
			err := rover.Move()
			if err != nil {
				return "", nil, err
			}
			continue
		}

		direction := spatial.DirectionFromString(navigationCommand)
		if direction == spatial.DirectionUnknown {
			return "", nil, ErrParsingEnvironmentCommand(commands[1])
		}

		rover.ChangeHeading(direction)
	}

	currentPosition, err := rover.CurrentPosition()
	if err != nil {
		return "", nil, err
	}

	currentHeading := spatial.HeadingToString(rover.CurrentHeading())
	roverStats := fmt.Sprintf("%v %v %v", currentPosition.X, currentPosition.Y, currentHeading)

	if len(commands) == 1 {
		return roverStats, nil, nil
	}
	return roverStats, commands[1:], nil
}

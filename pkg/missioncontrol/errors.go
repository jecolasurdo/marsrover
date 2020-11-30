package missioncontrol

import "fmt"

// ErrParsingEnvironmentCommand occurs when an environment command is malformed.
func ErrParsingEnvironmentCommand(cmd string) error {
	return fmt.Errorf("error parsing environment command '%v'", cmd)
}

// ErrParsingRoverCommand is returned if the commands supplied to the rover
// are sufficient to control the rover.
func ErrParsingRoverCommand(cmd string) error {
	return fmt.Errorf("the supplied commands are insufficient to move a rover. commands: '%v'", cmd)
}

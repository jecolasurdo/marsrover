package missioncontrol_test

import "testing"

func Test_NewMission(t *testing.T) {
	t.Skip()
	// returns a mission
}

func Test_ExecutionMission(t *testing.T) {
	t.Skip()
	// happy path returns statuses, nil
	// error establishing environment returns error
	// no commands, returns nil, nil
	// invalid commands returns nil, error
}

func Test_EstablishEnvironment(t *testing.T) {
	t.Skip()
	// happy path returns environment, remaining commands, nil
	// no commands, returns nil, nil, error
	// not two coordinates returns error
	// invalid x returns error
	// invalid y returns error
	// newenvironment error returns error
	// one valid command returns env, nil, nil
}

func Test_DeployAndNavigateRover(t *testing.T) {
	t.Skip()
	// happy path returns status, remaining commands, nil
	// nil environment returns error
	// less than 2 commands returns error
	// valid placement invalid navigate returns error
	// invalid placement valid navigate returns error
}

func Test_PlaceRoverInEnvironment(t *testing.T) {
	t.Skip()
	// happy path returns rover, remaining commands, nil
	// nil environment returns error
	// less than three positions returns error
	// invalid x returns error
	// invalid y returns error
	// invalid heading returns error
	// error launching rover returns error
}

func Test_NavigateRover(t *testing.T) {
	t.Skip()
	// happy path returns stats, remaining commands, nil
	// nil rover returns error
	// empty commands returns stats, remaining commands, nil
	// move error returns error
	// currentposition error returns error
	// move command calls move
	// invalid direction returns error
}

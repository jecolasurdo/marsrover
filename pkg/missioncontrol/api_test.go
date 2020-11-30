package missioncontrol_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_environmentiface "github.com/jecolasurdo/marsrover/mocks/environment"
	mock_roveriface "github.com/jecolasurdo/marsrover/mocks/rover"
	"github.com/jecolasurdo/marsrover/pkg/environment"
	"github.com/jecolasurdo/marsrover/pkg/environment/environmentiface"
	"github.com/jecolasurdo/marsrover/pkg/missioncontrol"
	"github.com/jecolasurdo/marsrover/pkg/objects"
	"github.com/jecolasurdo/marsrover/pkg/objects/roveriface"
	"github.com/jecolasurdo/marsrover/pkg/spatial"
	"github.com/stretchr/testify/assert"
)

func Test_NewMission(t *testing.T) {
	t.Run("panic if envBuilder nil", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		assert.Panics(t, func() {
			missioncontrol.NewMission(nil, mock_roveriface.NewMockRoverBuilder(ctrl))
		})
	})

	t.Run("panic if roverBuilder nil", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		assert.Panics(t, func() {
			missioncontrol.NewMission(mock_environmentiface.NewMockEnvironmentBuilder(ctrl), nil)
		})
	})

	t.Run("happy path returns mission", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mission := missioncontrol.NewMission(
			mock_environmentiface.NewMockEnvironmentBuilder(ctrl),
			mock_roveriface.NewMockRoverBuilder(ctrl),
		)
		assert.NotNil(t, mission)
	})
}

func Test_ExecutionMission(t *testing.T) {
	testCases := []struct {
		name     string
		commands []string
		expStats []string
		expErr   error
	}{
		{
			name:     "happy path",
			commands: []string{"5 5", "1 2 N", "LMLMLMLMM", "3 3 E", "MMRMMRMRRM"},
			expStats: []string{"1 3 N", "5 1 E"},
			expErr:   nil,
		},
	}
	// happy path returns statuses, nil
	// error establishing environment returns error
	// no commands, returns nil, nil
	// invalid commands returns nil, error

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			roverBuilder := mock_roveriface.NewMockRoverBuilder(ctrl)
			roverBuilder.EXPECT().
				LaunchRover(gomock.Any(), gomock.Any(), gomock.Any()).
				AnyTimes().
				DoAndReturn(
					func(h spatial.Heading, p spatial.Point, env environmentiface.Environmenter) (roveriface.RoverAPI, error) {
						return objects.Rover{}.LaunchRover(h, p, env)
					})

			envBuilder := mock_environmentiface.NewMockEnvironmentBuilder(ctrl)
			envBuilder.EXPECT().
				NewEnvironment(gomock.Any()).
				AnyTimes().
				DoAndReturn(func(p spatial.Point) environmentiface.Environmenter {
					return environment.Plateau{}.NewPlateau(p)
				})

			mission := missioncontrol.NewMission(envBuilder, roverBuilder)
			stats, err := mission.ExecuteMission(testCase.commands)

			assert.Equal(t, testCase.expStats, stats)
			if testCase.expErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, testCase.expErr, err.Error())
			}
		})
	}
}

// func Test_EstablishEnvironment(t *testing.T) {
// 	t.Skip()
// 	// happy path returns environment, remaining commands, nil
// 	// no commands, returns nil, nil, error
// 	// not two coordinates returns error
// 	// invalid x returns error
// 	// invalid y returns error
// 	// newenvironment error returns error
// 	// one valid command returns env, nil, nil
// }

// func Test_DeployAndNavigateRover(t *testing.T) {
// 	t.Skip()
// 	// happy path returns status, remaining commands, nil
// 	// nil environment returns error
// 	// less than 2 commands returns error
// 	// valid placement invalid navigate returns error
// 	// invalid placement valid navigate returns error
// }

// func Test_PlaceRoverInEnvironment(t *testing.T) {
// 	t.Skip()
// 	// happy path returns rover, remaining commands, nil
// 	// nil environment returns error
// 	// less than three positions returns error
// 	// invalid x returns error
// 	// invalid y returns error
// 	// invalid heading returns error
// 	// error launching rover returns error
// }

// func Test_NavigateRover(t *testing.T) {
// 	t.Skip()
// 	// happy path returns stats, remaining commands, nil
// 	// nil rover returns error
// 	// empty commands returns stats, remaining commands, nil
// 	// move error returns error
// 	// currentposition error returns error
// 	// move command calls move
// 	// invalid direction returns error
// }

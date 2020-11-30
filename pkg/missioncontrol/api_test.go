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
			name:     "spec example 1",
			commands: []string{"5 5", "1 2 N", "LMLMLMLMM", "3 3 E", "MMRMMRMRRM"},
			expStats: []string{"1 3 N", "5 1 E"},
			expErr:   nil,
		},
		{
			name:     "spec example 2",
			commands: []string{"3 3", "0 0 S", "LMMLM", "1 2 W", "LMLMRM"},
			expStats: []string{"2 1 N", "1 0 S"},
			expErr:   nil,
		},
		{
			name:     "no commands returns nil, nil",
			commands: nil,
			expStats: []string{},
			expErr:   nil,
		},
		{
			name:     "invalid environment command",
			commands: []string{"invalid"},
			expStats: nil,
			expErr:   missioncontrol.ErrParsingEnvironmentCommand("invalid"),
		},
		{
			name:     "invalid env x",
			commands: []string{"a 10"},
			expStats: nil,
			expErr:   missioncontrol.ErrParsingEnvironmentCommand("a 10"),
		},
		{
			name:     "invalid env y",
			commands: []string{"10 a"},
			expStats: nil,
			expErr:   missioncontrol.ErrParsingEnvironmentCommand("10 a"),
		},
		{
			name:     "incomplete rover command",
			commands: []string{"10 10", "1 2"},
			expStats: nil,
			expErr:   missioncontrol.ErrParsingRoverCommand("expected at least two commands"),
		},
		{
			name:     "invalid position command",
			commands: []string{"10 10", "1 2 3 4", "M"},
			expStats: nil,
			expErr:   missioncontrol.ErrParsingRoverCommand("1 2 3 4"),
		},
		{
			name:     "invalid rover x",
			commands: []string{"10 10", "a 2 N", "M"},
			expStats: nil,
			expErr:   missioncontrol.ErrParsingRoverCommand("a 2 N"),
		},
		{
			name:     "invalid rover y",
			commands: []string{"10 10", "1 a N", "M"},
			expStats: nil,
			expErr:   missioncontrol.ErrParsingRoverCommand("1 a N"),
		},
		{
			name:     "invalid heading",
			commands: []string{"10 10", "1 2 F", "M"},
			expStats: nil,
			expErr:   missioncontrol.ErrParsingRoverCommand("1 2 F"),
		},
		{
			name:     "elided movement results in an error",
			commands: []string{"10 10", "1 2 F", ""},
			expStats: nil,
			expErr:   missioncontrol.ErrParsingRoverCommand("1 2 F"),
		},
		{
			name:     "invalid movement cmd returns error",
			commands: []string{"10 10", "1 2 N", "D"},
			expStats: nil,
			expErr:   missioncontrol.ErrParsingRoverCommand("D"),
		},
	}

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
				assert.EqualError(t, err, testCase.expErr.Error())
			}
		})
	}
}

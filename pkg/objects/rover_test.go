package objects_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	mock_environmentiface "github.com/jecolasurdo/marsrover/mocks/environment"
	"github.com/jecolasurdo/marsrover/pkg/environment/environmenttypes"
	"github.com/jecolasurdo/marsrover/pkg/objects"
	"github.com/jecolasurdo/marsrover/pkg/spatial"
	"github.com/stretchr/testify/assert"
)

func Test_LaunchRover(t *testing.T) {
	t.Run("launching a valid rover returns a properly initialized rover", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		env := mock_environmentiface.NewMockEnvironmenter(ctrl)
		env.EXPECT().
			PlaceObject(gomock.Any(), gomock.Any()).
			Return(nil).
			AnyTimes()

		rover, err := objects.LaunchRover(spatial.HeadingNorth, spatial.NewPoint(1, 1), env)
		assert.Nil(t, err)
		assert.NotNil(t, rover)
	})

	t.Run("launching into an illegal position returns nil, error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		testError := "test error"
		env := mock_environmentiface.NewMockEnvironmenter(ctrl)
		env.EXPECT().
			PlaceObject(gomock.Any(), gomock.Any()).
			Return(fmt.Errorf(testError)).
			AnyTimes()

		rover, err := objects.LaunchRover(spatial.HeadingNorth, spatial.NewPoint(1, 1), env)
		assert.Nil(t, rover)
		assert.EqualError(t, err, testError)
	})

	t.Run("launching into a position already occupied by another rover returns nil, error", func(t *testing.T) {
		t.Fatal("not implemented")
	})
}

func Test_RoverCurrentPosition(t *testing.T) {
	t.Run("the rover's position is reported so long as it is still within its environment", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		env := mock_environmentiface.NewMockEnvironmenter(ctrl)
		env.EXPECT().
			PlaceObject(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		initialPosition := spatial.NewPoint(1, 1)
		env.EXPECT().
			FindObject(gomock.Any()).
			Return(true, &environmenttypes.ObjectPosition{Position: initialPosition}).
			Times(1)

		rover, err := objects.LaunchRover(spatial.HeadingNorth, initialPosition, env)
		assert.Nil(t, err)

		currentPosition, err := rover.CurrentPosition()
		assert.Nil(t, err)
		assert.Equal(t, initialPosition, *currentPosition)
	})

	t.Run("an error is returned if rover is no longer recognized by its environment", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		env := mock_environmentiface.NewMockEnvironmenter(ctrl)
		env.EXPECT().
			PlaceObject(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		initialPosition := spatial.NewPoint(1, 1)
		env.EXPECT().
			FindObject(gomock.Any()).
			Return(false, nil).
			Times(1)

		rover, err := objects.LaunchRover(spatial.HeadingNorth, initialPosition, env)
		assert.Nil(t, err)

		currentPosition, err := rover.CurrentPosition()
		assert.Nil(t, currentPosition)
		assert.EqualError(t, err, "this rover no longer exists within its environment")
	})
}

func Test_RoverCurrentHeading(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	env := mock_environmentiface.NewMockEnvironmenter(ctrl)
	env.EXPECT().
		PlaceObject(gomock.Any(), gomock.Any()).
		Return(nil).
		Times(1)

	initialPosition := spatial.NewPoint(1, 1)
	rover, err := objects.LaunchRover(spatial.HeadingNorth, initialPosition, env)
	assert.Nil(t, err)

	currentHeading := rover.CurrentHeading()
	assert.Equal(t, spatial.HeadingNorth, currentHeading)
}

func Test_RoverChangeHeading(t *testing.T) {
	testCases := []struct{ initialHeading, direction, resultingHeading string }{
		{"N", "R", "E"},
		{"E", "R", "S"},
		{"S", "R", "W"},
		{"W", "R", "N"},
		{"N", "L", "W"},
		{"W", "L", "S"},
		{"S", "L", "E"},
		{"E", "L", "N"},
	}

	for _, testCase := range testCases {
		testName := testCase.initialHeading + testCase.direction + testCase.resultingHeading
		t.Run(testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			env := mock_environmentiface.NewMockEnvironmenter(ctrl)
			env.EXPECT().
				PlaceObject(gomock.Any(), gomock.Any()).
				Return(nil).
				AnyTimes()

			initialPosition := spatial.NewPoint(1, 1)
			initialHeading := spatial.HeadingFromString(testCase.initialHeading)
			rover, err := objects.LaunchRover(initialHeading, initialPosition, env)
			assert.Nil(t, err)

			rover.ChangeHeading(spatial.DirectionFromString(testCase.direction))
			newHeading := rover.CurrentHeading()
			expectedHeading := spatial.HeadingFromString(testCase.resultingHeading)
			assert.Equal(t, expectedHeading, newHeading)
		})
	}
}

func Test_RoverMove(t *testing.T) {
	t.Run("move in each direction succeeds in proper calls to environment", func(t *testing.T) {
		testCases := []struct {
			initialHeading    spatial.Heading
			initialPosition   spatial.Point
			resultingPosition spatial.Point
		}{
			{
				spatial.HeadingNorth,
				spatial.NewPoint(3, 7),
				spatial.NewPoint(3, 8),
			},
			{
				spatial.HeadingEast,
				spatial.NewPoint(3, 7),
				spatial.NewPoint(4, 7),
			},
			{
				spatial.HeadingSouth,
				spatial.NewPoint(3, 7),
				spatial.NewPoint(3, 6),
			},
			{
				spatial.HeadingWest,
				spatial.NewPoint(3, 7),
				spatial.NewPoint(2, 7),
			},
		}

		for i, testCase := range testCases {
			testName := fmt.Sprintf("test case %v", i)
			t.Run(testName, func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				env := mock_environmentiface.NewMockEnvironmenter(ctrl)
				env.EXPECT().
					PlaceObject(gomock.Any(), gomock.Any()).
					Return(nil).
					Times(1)

				rover, err := objects.LaunchRover(testCase.initialHeading, testCase.initialPosition, env)
				assert.Nil(t, err)

				env.EXPECT().
					FindObject(rover).
					Return(true, &environmenttypes.ObjectPosition{
						Object:   rover,
						Position: testCase.initialPosition,
					}).
					Times(1)

				env.EXPECT().
					RecordMovement(rover, testCase.resultingPosition).
					Return(nil).
					Times(1)

				err = rover.Move()
				assert.Nil(t, err)
			})
		}
	})

	t.Run("move reports error if the object no longer exists in the environment", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		env := mock_environmentiface.NewMockEnvironmenter(ctrl)
		env.EXPECT().
			PlaceObject(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		rover, err := objects.LaunchRover(spatial.HeadingNorth, spatial.NewPoint(5, 5), env)
		assert.Nil(t, err)

		env.EXPECT().
			FindObject(rover).
			Return(false, nil).
			Times(1)

		env.EXPECT().
			RecordMovement(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(0)

		err = rover.Move()
		assert.EqualError(t, err, "this rover no longer exists within its environment")
	})

	t.Run("move reports error if there's an error recording the movement in the environment",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			env := mock_environmentiface.NewMockEnvironmenter(ctrl)
			env.EXPECT().
				PlaceObject(gomock.Any(), gomock.Any()).
				Return(nil).
				Times(1)

			initialPosition := spatial.NewPoint(5, 5)
			rover, err := objects.LaunchRover(spatial.HeadingNorth, initialPosition, env)
			assert.Nil(t, err)

			env.EXPECT().
				FindObject(rover).
				Return(true, &environmenttypes.ObjectPosition{
					Object:   rover,
					Position: initialPosition,
				}).
				Times(1)

			testError := fmt.Errorf("test error")
			env.EXPECT().
				RecordMovement(gomock.Any(), gomock.Any()).
				Return(testError).
				Times(1)

			err = rover.Move()
			assert.EqualError(t, err, testError.Error())
		})

	t.Run("moving into occupied space returns an error and does not call record movement",
		func(t *testing.T) {
			t.Fatal("not implemented")
			// ctrl := gomock.NewController(t)
			// defer ctrl.Finish()

			// env := mock_environmentiface.NewMockEnvironmenter(ctrl)
			// env.EXPECT().
			// 	PlaceObject(gomock.Any(), gomock.Any()).
			// 	Return(nil).
			// 	Times(2)

			// roverBPosition := spatial.NewPoint(4, 5)
			// roverB, err := objects.LaunchRover(spatial.HeadingEast, roverBPosition, env)
			// assert.Nil(t, err)

			// err = roverB.Move()
			// assert.EqualError(t, err, "unable to move, another object detectected at destination position")
		})
}

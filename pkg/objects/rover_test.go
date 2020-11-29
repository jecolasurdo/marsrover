package objects_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	mock_environmentiface "github.com/jecolasurdo/marsrover/mocks/environment"
	"github.com/jecolasurdo/marsrover/pkg/coordinate"
	"github.com/jecolasurdo/marsrover/pkg/environment/environmenttypes"
	"github.com/jecolasurdo/marsrover/pkg/objects"
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

		rover, err := objects.LaunchRover(objects.HeadingNorth, coordinate.NewPoint(1, 1), env)
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

		rover, err := objects.LaunchRover(objects.HeadingNorth, coordinate.NewPoint(1, 1), env)
		assert.Nil(t, rover)
		assert.EqualError(t, err, testError)
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

		initialPosition := coordinate.NewPoint(1, 1)
		env.EXPECT().
			FindObject(gomock.Any()).
			Return(true, &environmenttypes.ObjectPosition{Position: initialPosition}).
			Times(1)

		rover, err := objects.LaunchRover(objects.HeadingNorth, initialPosition, env)
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

		initialPosition := coordinate.NewPoint(1, 1)
		env.EXPECT().
			FindObject(gomock.Any()).
			Return(false, nil).
			Times(1)

		rover, err := objects.LaunchRover(objects.HeadingNorth, initialPosition, env)
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

	initialPosition := coordinate.NewPoint(1, 1)
	rover, err := objects.LaunchRover(objects.HeadingNorth, initialPosition, env)
	assert.Nil(t, err)

	currentHeading := rover.CurrentHeading()
	assert.Equal(t, objects.HeadingNorth, currentHeading)
}

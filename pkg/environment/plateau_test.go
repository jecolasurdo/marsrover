package environment_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_objectiface "github.com/jecolasurdo/marsrover/mocks/object"
	"github.com/jecolasurdo/marsrover/pkg/coordinate"
	"github.com/jecolasurdo/marsrover/pkg/environment"
	"github.com/jecolasurdo/marsrover/pkg/object/objectiface"
)

func Test_PlateauGetDimensions(t *testing.T) {
	expectedDimensions := coordinate.Point{
		X: 5,
		Y: 9,
	}
	p := environment.NewPlateau(expectedDimensions)

	if p.GetDimensions() != expectedDimensions {
		t.Fail()
	}
}

func Test_PlateauShowObjects(t *testing.T) {
	t.Run("returns empty map if no objects present", func(t *testing.T) {
		p := environment.NewPlateau(coordinate.Point{X: 10, Y: 10})

		actResult := p.ShowObjects()
		expResult := make(map[coordinate.Point][]objectiface.Objecter)

		assert.Equal(t, expResult, actResult)
	})

	t.Run("returns object map if objects exist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockObject1 := mock_objectiface.NewMockObjecter(ctrl)
		mockObject1.EXPECT().ID().Return("1").AnyTimes()

		mockObject2 := mock_objectiface.NewMockObjecter(ctrl)
		mockObject2.EXPECT().ID().Return("2").AnyTimes()

		mockObjects := map[coordinate.Point][]objectiface.Objecter{
			{X: 1, Y: 2}: {mockObject1},
			{X: 2, Y: 3}: {mockObject2},
		}

		p := environment.NewPlateau(coordinate.Point{X: 10, Y: 10})
		for position, objects := range mockObjects {
			for _, object := range objects {
				err := p.PlaceObject(object, position)
				if err != nil {
					panic(err)
				}
			}
		}

		actResult := p.ShowObjects()

		assert.Equal(t, mockObjects, actResult)
	})
}

func Test_PlateauPlaceObjects(t *testing.T) {
	t.Run("nil object returns an error", func(t *testing.T) {
		p := environment.NewPlateau(coordinate.Point{X: 10, Y: 10})
		err := p.PlaceObject(nil, coordinate.Point{X: 1, Y: 1})
		assert.EqualError(t, err, "a nil object cannot be placed in the environment")
	})

	t.Run("illegal coordinate returns an error", func(t *testing.T) {
		testCases := []struct {
			Name string
			X    int
			Y    int
		}{
			{
				Name: "above upper bound",
				X:    11,
				Y:    11,
			},
			{
				Name: "below lower bound",
				X:    -1,
				Y:    -1,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.Name, func(t *testing.T) {
				ctrl := gomock.NewController(t)
				mockObject := mock_objectiface.NewMockObjecter(ctrl)
				defer ctrl.Finish()

				p := environment.NewPlateau(coordinate.Point{X: 10, Y: 10})
				err := p.PlaceObject(mockObject, coordinate.Point{X: testCase.X, Y: testCase.Y})
				assert.EqualError(t, err, "an object cannot be placed outside the bounds of the environment")
			})
		}
	})

	// object already in environment returns an error (same object different coordinate)
	// object arleady in env... (same objects sharing coordinate)
	// placing an object places the object
	// multiple objects can be placed at the same location
}

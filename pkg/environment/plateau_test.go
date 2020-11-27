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
				Name: "X above upper",
				X:    11,
				Y:    9,
			},
			{
				Name: "Y above upper",
				X:    9,
				Y:    11,
			},
			{
				Name: "x below lower",
				X:    -1,
				Y:    1,
			},
			{
				Name: "y below lower",
				X:    1,
				Y:    -1,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.Name, func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				mockObject := mock_objectiface.NewMockObjecter(ctrl)

				p := environment.NewPlateau(coordinate.Point{X: 10, Y: 10})
				err := p.PlaceObject(mockObject, coordinate.Point{X: testCase.X, Y: testCase.Y})
				assert.EqualError(t, err, "an object cannot be placed outside the bounds of the environment")
			})
		}
	})

	t.Run("an object can only appear once in an environment", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockObject := mock_objectiface.NewMockObjecter(ctrl)
		mockObject.EXPECT().ID().Return("A").AnyTimes()

		p := environment.NewPlateau(coordinate.Point{X: 10, Y: 10})
		err := p.PlaceObject(mockObject, coordinate.Point{X: 1, Y: 1})
		assert.NoError(t, err)

		err = p.PlaceObject(mockObject, coordinate.Point{X: 1, Y: 1})
		assert.EqualError(t, err, "object with ID 'A' already exists within the environment")
	})

	t.Run("multiple objects can be placed at the same location", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockObjectA := mock_objectiface.NewMockObjecter(ctrl)
		mockObjectA.EXPECT().ID().Return("A").AnyTimes()

		mockObjectB := mock_objectiface.NewMockObjecter(ctrl)
		mockObjectB.EXPECT().ID().Return("B").AnyTimes()

		sharedLocation := coordinate.Point{X: 1, Y: 1}

		p := environment.NewPlateau(coordinate.Point{X: 10, Y: 10})

		err := p.PlaceObject(mockObjectA, sharedLocation)
		assert.NoError(t, err)

		err = p.PlaceObject(mockObjectB, sharedLocation)
		assert.NoError(t, err)
	})

	t.Run("multiple objects can be placed at different locations", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockObjectA := mock_objectiface.NewMockObjecter(ctrl)
		mockObjectA.EXPECT().ID().Return("A").AnyTimes()

		mockObjectB := mock_objectiface.NewMockObjecter(ctrl)
		mockObjectB.EXPECT().ID().Return("B").AnyTimes()

		p := environment.NewPlateau(coordinate.Point{X: 10, Y: 10})

		locationOne := coordinate.Point{X: 1, Y: 1}
		err := p.PlaceObject(mockObjectA, locationOne)
		assert.NoError(t, err)

		locationTwo := coordinate.Point{X: 2, Y: 2}
		err = p.PlaceObject(mockObjectB, locationTwo)
		assert.NoError(t, err)
	})
}

func Test_PlateauRecordMovement(t *testing.T) {
	// An object cannot be moved to an coordinate outside of the environment.
	// Moving an object effectively moves the object
	// Moving an object to a position where another object exists does not
	//     disturb the other object
}

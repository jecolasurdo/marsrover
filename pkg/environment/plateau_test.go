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

// We expect ShowObjects to return an empty map if there are no objects in the
// environment.
func Test_PlateauShowObjects_Empty(t *testing.T) {
	p := environment.NewPlateau(coordinate.Point{X: 10, Y: 10})

	actResult := p.ShowObjects()
	expResult := make(map[coordinate.Point][]objectiface.Objecter)

	assert.Equal(t, expResult, actResult)
}

// We expect ShowObjects to return a map of objects if objects have been added
// to the environment.
func Test_PlateauShowObjects_NotEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockObject1 := mock_objectiface.NewMockObjecter(ctrl)
	mockObject1.EXPECT().ID().Return("1").AnyTimes()

	mockObject2 := mock_objectiface.NewMockObjecter(ctrl)
	mockObject2.EXPECT().ID().Return("2").AnyTimes()

	mockObjects := map[coordinate.Point][]objectiface.Objecter{
		coordinate.Point{X: 1, Y: 2}: []objectiface.Objecter{mockObject1},
		coordinate.Point{X: 2, Y: 3}: []objectiface.Objecter{mockObject2},
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
}

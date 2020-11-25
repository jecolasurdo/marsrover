package environment_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

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
	testCases := []struct {
		name      string
		objects   map[coordinate.Point][]objectiface.Objecter
		expResult map[coordinate.Point][]objectiface.Objecter
	}{
		{
			name:      "empty environment returns empty map",
			objects:   make(map[coordinate.Point][]objectiface.Objecter),
			expResult: make(map[coordinate.Point][]objectiface.Objecter),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			p := environment.NewPlateau(coordinate.Point{X: 10, Y: 10})
			for position, objects := range testCase.objects {
				for _, object := range objects {
					err := p.PlaceObject(object, position)
					if err != nil {
						panic(err)
					}
				}
			}

			actResult := p.ShowObjects()

			assert.Equal(t, testCase.expResult, actResult)
		})
	}

}

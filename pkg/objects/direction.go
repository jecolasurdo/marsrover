package objects

// Direction is a general direction (left or right)
type Direction string

// Directions that can be applied to an object.
const (
	DirectionLeft  = "L"
	DirectionRight = "R"
)

// DirectionFromString converts a direction string ("L", "R") to a Direction.
// This function will panic if it receives an invalid direction string.
func DirectionFromString(d string) Direction {
	switch d {
	case "L":
		return DirectionLeft
	case "R":
		return DirectionRight
	default:
		panic("unknown direction")
	}
}

package spatial

// Direction is a general direction (left or right)
type Direction string

// Directions that can be applied to an object.
const (
	DirectionUnknown = ""
	DirectionLeft    = "L"
	DirectionRight   = "R"
)

// DirectionFromString converts a direction string ("L", "R") to a Direction.
// If the supplied value is not L or R, the function will return DirectionUnknown.
func DirectionFromString(d string) Direction {
	switch d {
	case "L":
		return DirectionLeft
	case "R":
		return DirectionRight
	default:
		return DirectionUnknown
	}
}

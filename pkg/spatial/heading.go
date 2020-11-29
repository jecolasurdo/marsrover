package spatial

// Heading represents the direction that an object may be pointing.
type Heading int

// Headings generally available for objects to use.
const (
	HeadingNorth Heading = 0
	HeadingEast  Heading = 1
	HeadingSouth Heading = 2
	HeadingWest  Heading = 3
)

// HeadingFromString converts a heading string ("N", "E", "S", "W") to a Heading value.
// This function will panic if it receives an invalid heading string.
func HeadingFromString(h string) Heading {
	switch h {
	case "N":
		return HeadingNorth
	case "E":
		return HeadingEast
	case "S":
		return HeadingSouth
	case "W":
		return HeadingWest
	default:
		panic("unknown heading")
	}
}

// Cardinals is an array of cardinal directions (Headings).
//
// Note: It is important that the indices of the values in this array match the
// respective value for the Heading constants.
var Cardinals = [4]Heading{HeadingNorth, HeadingEast, HeadingSouth, HeadingWest}

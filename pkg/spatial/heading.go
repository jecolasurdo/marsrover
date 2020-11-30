package spatial

// Heading represents the direction that an object may be pointing.
type Heading int

// Headings generally available for objects to use.
const (
	HeadingUnknown Heading = -1
	HeadingNorth   Heading = 0
	HeadingEast    Heading = 1
	HeadingSouth   Heading = 2
	HeadingWest    Heading = 3
)

// Cardinals is an array of cardinal directions (Headings).
//
// Note: It is important that the indices of the values in this array match the
// respective value for the Heading constants.
var Cardinals = [4]Heading{HeadingNorth, HeadingEast, HeadingSouth, HeadingWest}

// HeadingFromString converts a heading string ("N", "E", "S", "W") to a Heading value.
// If the supplied value cannot be mapped to a heading, this function will return
// HeadingUnknown
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
		return HeadingUnknown
	}
}

// HeadingToString converts a Heading to a string.
func HeadingToString(h Heading) string {
	switch h {
	case HeadingNorth:
		return "N"
	case HeadingEast:
		return "E"
	case HeadingSouth:
		return "S"
	case HeadingWest:
		return "W"
	default:
		return ""
	}
}

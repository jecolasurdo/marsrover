package objects

// Heading represents the direction that an object may be pointing.
type Heading string

// Headings generally available for objects to use.
const (
	HeadingNorth Heading = "N"
	HeadingEast  Heading = "E"
	HeadingSouth Heading = "S"
	HeadingWest  Heading = "W"
)

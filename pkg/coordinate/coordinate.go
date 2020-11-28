package coordinate

// Point is a coordinate in a two dimensional plane.
type Point struct {
	X int
	Y int
}

// NewPoint instantiates a new point.
func NewPoint(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

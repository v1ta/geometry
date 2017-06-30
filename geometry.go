package geometry

// Geometry defines a means of interacting with Rectangle elements in R^2.
type Geometry interface {
	Overlap(r *Rectangle) (bool, []Point)
	Contains(r *Rectangle) bool
	Adjacenct(r *Rectangle) bool
	EnumeratePoints(r *Rectangle) []Point
	IsWellFormed(r *Rectangle) bool
	ContainsCoords(p *Point) bool
	String() string
}

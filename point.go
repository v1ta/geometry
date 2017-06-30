package geometry

import "fmt"

// Point a vector in R^2
type Point struct {
	X float64
	Y float64
}

// toString for Point
func (p *Point) String() string {
	return fmt.Sprintf("<%.2f, %.2f>", p.X, p.Y)
}

// ContainsCoords checks if a coordinate from a point is present in a set of Points
func (p *Point) ContainsCoords(points []Point) bool {
	for _, v := range points {
		if p.X == v.X || p.Y == v.Y {
			return true
		}
	}

	return false
}

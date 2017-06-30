package geometry

import "fmt"

// Point a vector in R^2.
type Point struct {
	X, Y float64
}

func (p *Point) String() string {
	return fmt.Sprintf("<%.2f, %.2f>", p.X, p.Y)
}

// ContainsCoords checks if a coordinate (X or Y) from a Point is present in a set of Points.
func (p *Point) ContainsCoords(points []Point) bool {
	for _, v := range points {
		if p.X == v.X || p.Y == v.Y {
			return true
		}
	}

	return false
}

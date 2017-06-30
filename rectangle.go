package geometry

import (
	"fmt"
	"math"
)

// Rectangle represented by a bottom left and top right Point.
type Rectangle struct {
	A, B Point
}

func (r *Rectangle) String() string {
	return fmt.Sprintf("%s | %s", r.A.String(), r.B.String())
}

// EnumeratePoints return all corner Points of a Rectangle.
func (r *Rectangle) EnumeratePoints() []Point {
	return []Point{r.A, Point{r.B.X, r.A.Y}, Point{r.A.X, r.B.Y}, r.B}
}

// IsWellFormed checks if Point a < Point b.
func (r *Rectangle) IsWellFormed() bool {
	return r.A.X < r.B.X && r.A.Y < r.B.Y
}

// Overlap checks if two Rectangles overlap.
func (r *Rectangle) Overlap(r2 *Rectangle) (bool, []Point) {
	A := Point{math.Max(r.A.X, r2.A.X), math.Max(r.A.Y, r2.A.Y)}
	B := Point{math.Min(r.B.X, r2.B.X), math.Min(r.B.Y, r2.B.Y)}
	intersections := []Point{}

	// Check if intersecting rectangle is well formed, a <= b
	if A.X > B.X || A.Y > B.Y {
		return false, intersections
	} else if A.X == B.X && A.Y == B.Y {
		return true, append(intersections, A)
	}

	// Find the points which have an edge on each rectangle
	// This works by making sure a given point p has a coordinate from each
	// set P where P is the set of points from one of the two Rectangle args
	for _, point := range []Point{A, Point{B.X, A.Y}, Point{A.X, B.Y}, B} {
		if func(point Point, rPoints, r2Points []Point) bool {
			return point.ContainsCoords(rPoints) && point.ContainsCoords(r2Points)
		}(point, r.EnumeratePoints(), r2.EnumeratePoints()) {
			intersections = append(intersections, point)
		}
	}

	return true, intersections
}

// Contains checks if a rectangle contains another.
func (r *Rectangle) Contains(r2 *Rectangle) bool {
	if r.A.X > r2.A.X || r.B.X < r2.B.X {
		return false
	}

	if r.A.Y > r2.A.Y || r.B.Y < r2.B.Y {
		return false
	}

	return true
}

// Adjacenct checks if two rectangles shares an edge.
func (r *Rectangle) Adjacenct(r2 *Rectangle) bool {

	// Overlap returns false if the rectangles aren't touching
	if doesOverlap, _ := r.Overlap(r2); !doesOverlap {
		return false
	}

	xLines := make(map[float64][]float64)
	yLines := make(map[float64][]float64)
	points := append(r.EnumeratePoints(), r2.EnumeratePoints()...)

	// Create a list of all points along a line for the opposing axis
	for _, point := range points {

		if _, ok := xLines[point.X]; ok {
			continue
		} else {
			adjacentPoints := []float64{}
			for _, p := range points {
				if p.X == point.X {
					adjacentPoints = append(adjacentPoints, p.Y)
				}
			}
			xLines[point.X] = adjacentPoints
		}

		if _, ok := yLines[point.Y]; ok {
			continue
		} else {
			adjacentPoints := []float64{}
			for _, p := range points {
				if p.Y == point.Y {
					adjacentPoints = append(adjacentPoints, p.X)
				}
			}
			yLines[point.Y] = adjacentPoints
		}

	}

	// if any given line in R^1 has 4 points, then we know there is an adjacent side
	for _, values := range xLines {
		if len(values) >= 4 {
			return true
		}
	}

	for _, values := range yLines {
		if len(values) >= 4 {
			return true
		}
	}

	return false
}

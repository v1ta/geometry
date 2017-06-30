package geometry

import (
	"fmt"
	"math"
)

// Rectangle represented by a bottom left and top right point
type Rectangle struct {
	a, b Point
}

// toString for Rectangle
func (r *Rectangle) String() string {
	return fmt.Sprintf("%s | %s", r.a.String(), r.b.String())
}

// EnumeratePoints return all corners of a Rectangle
func (r *Rectangle) EnumeratePoints() []Point {
	return []Point{r.a, Point{r.b.X, r.a.Y}, Point{r.a.X, r.b.Y}, r.b}
}

// IsWellFormed checks if Point a < Point b
func (r *Rectangle) IsWellFormed() bool {
	return r.a.X < r.b.X && r.a.Y < r.b.Y
}

// Overlap checks if two rectangles (r and r2) overlap
func (r *Rectangle) Overlap(r2 *Rectangle) (bool, []Point) {
	a := Point{math.Max(r.a.X, r2.a.X), math.Max(r.a.Y, r2.a.Y)}
	b := Point{math.Min(r.b.X, r2.b.X), math.Min(r.b.Y, r2.b.Y)}
	intersections := []Point{}

	// Check if intersecting rectangle is well formed, a <= b
	if a.X > b.X || a.Y > b.Y {
		return false, intersections
	} else if a.X == b.X && a.Y == b.Y {
		return true, append(intersections, a)
	}

	// Find the points which have an edge on each rectangle
	// This works by making sure a given point p has a coordinate from each
	// set P where P is the set of points from one of the two Rectangle args
	for _, point := range []Point{a, Point{b.X, a.Y}, Point{a.X, b.Y}, b} {
		if func(point Point, rPoints, r2Points []Point) bool {
			return point.ContainsCoords(rPoints) && point.ContainsCoords(r2Points)
		}(point, r.EnumeratePoints(), r2.EnumeratePoints()) {
			intersections = append(intersections, point)
		}
	}

	return true, intersections
}

// Contains checks if a rectangle (r) contains another (r2)
func (r *Rectangle) Contains(r2 *Rectangle) bool {
	if r.a.X > r2.a.X || r.b.X < r2.b.X {
		return false
	}

	if r.a.Y > r2.a.Y || r.b.Y < r2.b.Y {
		return false
	}

	return true
}

// Adjacenct checks if two rectangles shares an edge
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

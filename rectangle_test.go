package geometry

import "testing"

func ExecuteTest(expected, actual bool, t *testing.T) {
	if actual != expected {
		t.Errorf("\nExpected: %t\nActual: %t\n", expected, actual)
	}
}

func contains(points []Point, p Point) bool {
	for _, point := range points {
		if p == point {
			return true
		}
	}
	return false
}

func checkIntersections(intersections, validIntersections []Point, t *testing.T) {
	for _, intersection := range intersections {
		ExecuteTest(true, contains(validIntersections, intersection), t)
	}
}

// TestIsWellFormed validates that a rectangle can determine if a < b for Points a, b
func TestIsWellFormed(t *testing.T) {

	// Valid rectangles
	tc1 := Rectangle{Point{1, 1}, Point{2, 2}}
	tc2 := Rectangle{Point{-1, -1}, Point{2, 2}}

	ExecuteTest(true, tc1.IsWellFormed(), t)
	ExecuteTest(true, tc2.IsWellFormed(), t)

	// Invalid recangles
	tc3 := Rectangle{Point{4, 4}, Point{2, 2}}
	tc4 := Rectangle{Point{-10.213, 3}, Point{-123, 234}}

	ExecuteTest(false, tc3.IsWellFormed(), t)
	ExecuteTest(false, tc4.IsWellFormed(), t)
}

// TestOverlap test when two rectangles do or don't overlap
func TestOverlap(t *testing.T) {

	// Basic overlap
	tc1 := Rectangle{Point{1, 1}, Point{3, 3}}
	tc2 := Rectangle{Point{2, 2}, Point{4, 4}}
	doesOverlap, intersections := tc1.Overlap(&tc2)
	validIntersections := []Point{Point{3, 2}, Point{2, 3}}

	ExecuteTest(true, doesOverlap, t)
	checkIntersections(intersections, validIntersections, t)

	doesOverlap, intersections = tc2.Overlap(&tc1)

	ExecuteTest(true, doesOverlap, t)
	checkIntersections(intersections, validIntersections, t)

	// Corner overlap
	tc3 := Rectangle{Point{3, 3}, Point{4.2345234, 400.23423}}
	doesOverlap, intersections = tc3.Overlap(&tc1)

	ExecuteTest(true, doesOverlap, t)
	checkIntersections(intersections, []Point{Point{3, 3}}, t)

	// No overlap
	tc4 := Rectangle{Point{23452345, 123423.23423}, Point{999999999999, 999999999999999}}

	doesOverlap, intersections = tc4.Overlap(&tc3)

	ExecuteTest(false, doesOverlap, t)
	checkIntersections(intersections, []Point{}, t)

}

// TestContains test if one rectangle contains another
func TestContains(t *testing.T) {

	// Basic containment
	tc1 := Rectangle{Point{1, 1}, Point{10, 10}}
	tc2 := Rectangle{Point{2, 2}, Point{5, 5}}

	ExecuteTest(true, tc1.Contains(&tc2), t)
	ExecuteTest(false, tc2.Contains(&tc1), t)
	ExecuteTest(true, tc1.Contains(&tc1), t)

	// Overlap, but no containment
	tc3 := Rectangle{Point{2, 2}, Point{8, 8}}
	tc4 := Rectangle{Point{3, 3}, Point{7, 10}}

	ExecuteTest(false, tc3.Contains(&tc4), t)
	ExecuteTest(false, tc4.Contains(&tc3), t)

	// No overlap or containment
	tc5 := Rectangle{Point{20, 20}, Point{30, 30}}

	ExecuteTest(false, tc1.Contains(&tc5), t)
	ExecuteTest(false, tc5.Contains(&tc1), t)
}

// TestAdjacent check if two rectangles are adjacent
func TestAdjacent(t *testing.T) {

	// Basic adjaceny
	tc1 := Rectangle{Point{1, 1}, Point{4, 4}}
	tc2 := Rectangle{Point{4, 0}, Point{8, 8}}

	ExecuteTest(true, tc1.Adjacenct(&tc2), t)
	ExecuteTest(true, tc2.Adjacenct(&tc1), t)

	// Corner adjacency
	tc3 := Rectangle{Point{4, 4}, Point{8, 8}}

	ExecuteTest(true, tc1.Adjacenct(&tc3), t)
	ExecuteTest(true, tc3.Adjacenct(&tc1), t)

	// No adjacency
	tc4 := Rectangle{Point{20, 8}, Point{30, 30}}

	ExecuteTest(false, tc3.Adjacenct(&tc4), t)
	ExecuteTest(false, tc4.Adjacenct(&tc3), t)
}

// TestEnumeratePoints check if the correct corners are created when enumerating over a rectangle
func TestEnumeratePoints(t *testing.T) {

	tc1 := Rectangle{Point{1, 1}, Point{3, 3}}
	tc2 := Rectangle{Point{-1, -1}, Point{4, 4}}

	tc1Points := []Point{Point{1, 1}, Point{1, 3}, Point{1, 3}, Point{3, 3}}
	tc2Points := []Point{Point{-1, -1}, Point{-1, 4}, Point{4, -1}, Point{4, 4}}

	for _, point := range tc1Points {
		ExecuteTest(true, contains(tc1.EnumeratePoints(), point), t)
	}

	for _, point := range tc2Points {
		ExecuteTest(true, contains(tc2.EnumeratePoints(), point), t)
	}
}

// TestString test toString for Rectangle
func TestRectangleString(t *testing.T) {

	tc1 := Rectangle{Point{1, 1}, Point{2, 2}}
	tc1String := "<1.00, 1.00> | <2.00, 2.00>"

	ExecuteTest(true, tc1.String() == tc1String, t)
}

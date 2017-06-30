package geometry

import "testing"

// TestContainsCoords check whether a point has a coordinate present in a []Point.
func TestContainsCoords(t *testing.T) {

	pointArray := []Point{Point{-1, 4}, Point{5, 6}, Point{90.4, 2}}

	// Test for point which exists in list
	tc1 := Point{1, 2}
	tc2 := Point{-1, -2}

	ExecuteTest(true, tc1.ContainsCoords(pointArray), t)
	ExecuteTest(true, tc2.ContainsCoords(pointArray), t)

	// Test for points which don't exist in list
	tc3 := Point{0.2, 91}
	tc4 := Point{4, -1}

	ExecuteTest(false, tc3.ContainsCoords(pointArray), t)
	ExecuteTest(false, tc4.ContainsCoords(pointArray), t)

}

// TestPointString test toString for Point.
func TestPointString(t *testing.T) {

	tc1 := Point{5.23, 3.33}
	tc1String := "<5.23, 3.33>"

	ExecuteTest(true, tc1.String() == tc1String, t)

}

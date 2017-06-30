package main

import (
	"fmt"

	"github.com/v1ta/geometry"
)

func main() {

	fmt.Print("********** v1ta/geometery example **********\n\n")

	// Declare two rectangles
	r1 := geometry.Rectangle{A: geometry.Point{X: 2, Y: 1}, B: geometry.Point{X: 4, Y: 6}}
	r2 := geometry.Rectangle{A: geometry.Point{X: 1, Y: 3}, B: geometry.Point{X: 3, Y: 5}}

	// Print rectangles
	fmt.Printf("r1: %s\nr2: %s\n", r1.String(), r2.String())

	// Check for intersection, if it exists, print the intersecting points
	if overlap, intersections := r1.Overlap(&r2); overlap {
		for _, intersection := range intersections {
			fmt.Printf("intersecting point: %s\n", intersection.String())
		}
	}

	// Check for containment
	if contained := r1.Contains(&r2); contained {
		fmt.Print("r1 contains r2\n")
	} else {
		fmt.Print("r1 doesn't contain r2\n")
	}

	// Check for adjacency
	if adjacent := r1.Adjacenct(&r2); adjacent {
		fmt.Print("r1 is adjacent to r2\n")
	} else {
		fmt.Print("r1 is not adjacent to r2\n")
	}

}

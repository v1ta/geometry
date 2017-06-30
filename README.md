# Geometry
[![GoDoc](https://godoc.org/github.com/v1ta/geometry?status.svg)](https://godoc.org/github.com/v1ta/geometry)

### Usage
```
git clone https://github.com/v1ta/geometry.git $GOPATH/src/github.com/v1ta
go build
go test
go run cmd/main.go
```

### Code samples

### Create a point
```
p := geometry.Point{X: 1, Y: 2}
```

### Create a rectangle
```
r1 := geometry.Rectangle{A: geometry.Point{X: 2, Y: 1}, B: geometry.Point{X: 4, Y: 6}}
```

### Check for an intersection(s)
```
if overlap, intersections := r1.Overlap(&r2); overlap {
    for _, intersection := range intersections {
        fmt.Printf("intersecting point: %s\n", intersection.String())
    }
}
```

### Check for containment
```
if contained := r1.Contains(&r2); contained {
    fmt.Print("r1 contains r2\n")
} else {
    fmt.Print("r1 doesn't contain r2\n")
}
```

### Check for adjacency
```
if adjacent := r1.Adjacenct(&r2); adjacent {
    fmt.Print("r1 is adjacent to r2\n")
} else {
    fmt.Print("r1 is not adjacent to r2\n")
}
```
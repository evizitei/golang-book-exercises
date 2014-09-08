package main

import "math"

type Point struct {
	x float64
	y float64
}

func (point *Point) distanceFrom(other *Point) float64 {
	a := math.Abs(point.x - other.x)
	b := math.Abs(point.y - other.y)
	return math.Sqrt(a*a + b*b)
}

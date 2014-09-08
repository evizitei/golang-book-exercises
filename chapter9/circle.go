package main

import "math"

type Circle struct {
	origin *Point
	radius float64
}

func (circle *Circle) area() float64 {
	rSquared := circle.radius * circle.radius
	return math.Pi * rSquared
}

func (circle *Circle) perimeter() float64 {
	twoR := circle.radius + circle.radius
	return math.Pi * twoR
}

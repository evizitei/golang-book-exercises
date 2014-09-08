package main

import "math"

type Rectangle struct {
	origin   *Point
	opposite *Point
}

func (rect *Rectangle) width() float64 {
	return math.Abs(rect.origin.x - rect.opposite.x)
}

func (rect *Rectangle) height() float64 {
	return math.Abs(rect.origin.y - rect.opposite.y)
}

func (rect *Rectangle) area() float64 {
	return rect.width() * rect.height()
}

func (rect *Rectangle) perimeter() float64 {
	return rect.width()*2 + rect.height()*2
}

package main

type Rectangle struct {
	origin   *Point
	opposite *Point
}

func (rect *Rectangle) width() float64 {
	return rect.origin.x - rect.opposite.x
}

func (rect *Rectangle) height() float64 {
	return rect.origin.y - rect.opposite.y
}

func (rect *Rectangle) area() float64 {
	return rect.width() * rect.height()
}

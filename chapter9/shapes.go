package main

type Shape interface {
	area() float64
	perimeter() float64
}

type ShapeSet struct {
	shapes []Shape
}

func (set *ShapeSet) area() float64 {
	var area float64
	for _, s := range set.shapes {
		area += s.area()
	}
	return area
}

func (set *ShapeSet) perimeter() float64 {
	var perimeter float64
	for _, s := range set.shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

package main

import "fmt"

func main() {
	circle := Circle{origin: &Point{x: 0, y: 0}, radius: 5}
	rectangle := Rectangle{origin: &Point{x: 0, y: 0}, opposite: &Point{x: 10, y: 10}}
	set := ShapeSet{shapes: []Shape{&rectangle, &circle}}

	fmt.Println("Areas")
	fmt.Println(rectangle.area())
	fmt.Println(circle.area())
	fmt.Println(set.area())

	fmt.Println("Perimeters")
	fmt.Println(rectangle.perimeter())
	fmt.Println(circle.perimeter())
	fmt.Println(set.perimeter())
}

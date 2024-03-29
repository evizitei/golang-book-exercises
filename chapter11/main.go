package main

import "fmt"
import "golang-book/chapter11/math"

func main() {
	xs := []float64{1, 2, 3, 4}

	avg := math.Average(xs)
	fmt.Println(avg)

	max := math.Max(xs)
	fmt.Println(max)

	min := math.Min(xs)
	fmt.Println(min)
}

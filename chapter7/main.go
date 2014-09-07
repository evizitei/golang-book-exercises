package main

import "fmt"

func sum(xs ...float64) (total float64) {
	total = 0.0
	for _, v := range xs {
		total += v
	}
	return
}

func average(xs ...float64) float64 {
	total := sum(xs...)
	return total / float64(len(xs))
}

func main() {
	fmt.Println(average(98, 93, 77, 82, 83))
}

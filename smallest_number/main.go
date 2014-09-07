package main

import "fmt"

func findMin(numbers []int) int {
	min := numbers[0]
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}
	return min
}

func main() {
	numbers := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	min := findMin(numbers[:])
	fmt.Println("Min is: ", min)
}

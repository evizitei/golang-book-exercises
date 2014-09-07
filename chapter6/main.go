package main

import "fmt"

func averageOf(values []float64) float64 {
	var total float64 = 0
	for _, value := range values {
		total += value
	}
	return total / float64(len(values))
}

func findAverage() float64 {
	x := []float64{98, 93, 77, 82, 83}
	return averageOf(x[:])
}

func playWithAppend() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

func playWithCopy() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

func playWithMaps() {
	xMap := make(map[string]int)
	xMap["key"] = 10
	fmt.Println(xMap)

	intMap := make(map[int]int)
	intMap[5] = 42
	fmt.Println(intMap)
	delete(intMap, 5)
	fmt.Println(intMap)
}

func main() {
	fmt.Println("Average: ", findAverage())
	playWithAppend()
	playWithCopy()
	playWithMaps()
}

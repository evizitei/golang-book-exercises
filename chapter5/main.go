package main

import "fmt"

func evenOrOdd(x int) string {
	if x%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i, evenOrOdd(i))
	}
}

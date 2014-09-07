package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(x int) string {
	if x%3 == 0 {
		if x%5 == 0 {
			return "FizzBuzz"
		}
		return "Fizz"
	} else if x%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(x)
}

func main() {
	for i := 1; i <= 100; i++ {
		message := fizzBuzz(i)
		fmt.Println(message)
	}
}

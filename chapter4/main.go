package main

import "fmt"

func double(x float64) float64 {
	result := x * 2
	return result
}

func getNumberFromUser() float64 {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	return input
}

func giveUserResult(result float64) {
	fmt.Println(result)
}

func main() {
	input := getNumberFromUser()
	output := double(input)
	giveUserResult(output)
}

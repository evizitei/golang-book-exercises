package main

import "fmt"

func powersOf(base int) func() int {
	cursor := 1
	powerBuilder := func() (ret int) {
		ret = cursor
		cursor = cursor * base
		return
	}
	return powerBuilder
}

func performManyTimes(count int, operation func()) {
	for i := 0; i < count; i++ {
		operation()
	}
}

func getUserBase() int {
	fmt.Print("Enter the base: ")
	var base int
	fmt.Scanf("%d", &base)
	return base
}

func getUserPower() int {
	fmt.Print("Enter the power: ")
	var power int
	fmt.Scanf("%d", &power)
	return power
}

func main() {
	base := getUserBase()
	power := getUserPower()
	nextPowerOfX := powersOf(base)
	powerPrinter := func() {
		fmt.Println(nextPowerOfX())
	}
	performManyTimes(power+1, powerPrinter)
}

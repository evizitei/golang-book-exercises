package main

import (
	"fmt"
	"time"
)

func sendOne(c chan<- string) {
	for {
		c <- "from 1"
		time.Sleep(time.Second * 2)
	}
}

func sendTwo(c chan<- string) {
	for {
		c <- "from 2"
		time.Sleep(time.Second * 3)
	}
}

func pickPrinter(c1 chan string, c2 chan string) {
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go sendOne(c1)
	go sendTwo(c2)
	go pickPrinter(c1, c2)

	var input string
	fmt.Scanln(&input)
}

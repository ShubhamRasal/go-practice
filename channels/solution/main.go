package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("welcome to channels solution")

	channel := make(chan string)

	fmt.Println("first task: crush coffee")

	coffee := "crush"

	go func(crush string) {

		time.Sleep(5 * time.Second)
		fmt.Println("second task: brew expresso")

		time.Sleep(5 * time.Second)
		time.Sleep(5 * time.Second)
		time.Sleep(5 * time.Second)

		channel <- "expresso"

		fmt.Println("debug: send expresso")

	}(coffee)

	fmt.Println("third task: steam milk")

	// cup mdhe copfee otane aahe
	blackCoffee := <-channel

	fmt.Println("finally cofee made with", blackCoffee)

}

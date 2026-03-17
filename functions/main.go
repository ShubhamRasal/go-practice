package main

import (
	"fmt"
)

// checkPort checks if given input is port or not

func checkPort(port int) (string, error) {

	if port <= 0 || port >= 65535 {

		return "", fmt.Errorf("this is invalid port: %d", port)

	} else {

		return fmt.Sprintf("port %d is valid", port), nil
	}

}

func main() {
	// create and assign variable
	validPort := 8080

	msg, err := checkPort(validPort)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(msg)

	// wrong port case
	wrongPort := -1
	wrongmsg, wrongerr := checkPort(wrongPort)
	if wrongerr != nil {
		fmt.Println("Error:", wrongerr)
		return
	}
	fmt.Println(wrongPort, wrongmsg, wrongerr)

}

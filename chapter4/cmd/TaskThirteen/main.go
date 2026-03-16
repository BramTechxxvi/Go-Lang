package main

import "fmt"

func main() {

	var userInput int
	for {
		fmt.Println("Enter")
		fmt.Scan(&userInput)
		if userInput == 0 {
			break
		}
		fmt.Println(userInput)
	}

	factorial := 1
	for count := 1; count < 5; count++ {
		factorial *= count
	}
	fmt.Println(factorial)
}

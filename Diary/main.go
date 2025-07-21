package main

import "fmt"

func main() {

	printOut("Welcome to Diary App \n\n")

}

func printOut(message string) {
	fmt.Println(message)
}

func input(message string) string {
	for {
		fmt.Println(message)
		var response string
		_, err := fmt.Scanln(&response)
		if err == nil {
			return response
		}
		fmt.Println("Please input again, error")
	}
}

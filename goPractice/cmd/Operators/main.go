package main

import "fmt"

func main() {
	fmt.Print("Hello world \n")
	fmt.Println("Today is bright")

	//var content String = "I have a couple of bugs to fix"
	//fmt.Printf("%s\n", content)

	var firstNumber = 10
	var secondNumber = 121
	var thirdNumber = 390
	var fourthNumber = 987

	fmt.Println(addUp(firstNumber, secondNumber, thirdNumber, fourthNumber))
	fmt.Println(firstNumber - secondNumber)
	fmt.Printf("The result is: %d ", thirdNumber/30)
	fmt.Printf()
}

func addUp(first int, second int, rest ...int) int {
	var sum = first + second

	for _, value := range rest {
		sum += value
	}
	return sum
}

func substract(first int, second int, rest ...int) int {
	var result = first - second
	for _, value := range rest {
		result -= value
	}
	return result
}

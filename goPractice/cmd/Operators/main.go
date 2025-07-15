package main

import "fmt"

func main() {
	var firstNumber = 10
	var secondNumber = 121
	var thirdNumber = 390
	var fourthNumber = 987

	fmt.Printf("Sum is: %d\n\n", addUp(firstNumber, secondNumber, thirdNumber, fourthNumber))
	fmt.Printf("difference is: %d\n\n", subtract(thirdNumber, secondNumber))
	fmt.Printf("The result is: %d\n\n", thirdNumber/30)
	fmt.Printf("Printing 1 to %d\n", firstNumber)
	printNumbers(firstNumber)

}

func addUp(first int, second int, rest ...int) int {
	var sum = first + second
	for _, value := range rest {
		sum += value
	}
	return sum
}

func subtract(first int, second int, rest ...int) int {
	var result = first - second
	for _, value := range rest {
		result -= value
	}
	return result
}

func printNumbers(number int) {
	for count := 1; count <= number; count++ {
		fmt.Println(count)
	}
}

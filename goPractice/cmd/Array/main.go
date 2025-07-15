package main

import "fmt"

func main() {

	var intArray = []int{89, 0, 12}
	intArray[0] = 10
	evenNumbersInArray(intArray)
}

func evenNumbersInArray(numbers []int) {
	for count := 0; count < len(numbers); count++ {
		if numbers[count]%2 == 0 {
			fmt.Println(numbers[count])
		}
	}
}

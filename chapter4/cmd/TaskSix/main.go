package main

import "fmt"

func main() {

	var sum = 0
	var x = 1
	for x = 1; x <= 10; x++ {
		sum += x
	}
	fmt.Println(sum)
}

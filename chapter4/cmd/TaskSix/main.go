package main

import "fmt"

func main() {

	var sum = 0
	var x = 1
	for x <= 10 {
		sum += x
		x++
	}
	fmt.Println(sum)
}

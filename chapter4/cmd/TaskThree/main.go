package main

import "fmt"

func main() {

	var x = 5
	fmt.Println(x + 1)
	x = x + 1
	fmt.Println(x)
	x += 1
	fmt.Println(x)
	x++
	fmt.Println(x)
}

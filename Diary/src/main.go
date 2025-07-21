package src

import "fmt"

func main() {

}

func print(message string) {
	fmt.Println(message)
}

func input(message string) string {
	fmt.Println(message)
	fmt.Scanln(&message)
	return message
}

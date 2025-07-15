package main

import (
	"fmt"
	"strings"
)

func main() {

	var intArray = []int{89, 0, 12}
	intArray[0] = 10
	evenNumbersInArray(intArray)

	var words = []string{"Blade", "sky", "Honesty", "Crazy", "Money", "crypt"}
	wordsInArray(words)
}

func evenNumbersInArray(numbers []int) {
	for count := 0; count < len(numbers); count++ {
		if numbers[count]%2 == 0 {
			fmt.Println(numbers[count])
		}
	}
}

func wordsInArray(array []string) []string {
	var result []string
	for _, value := range array {
		if hasVowel(value) {
			result = append(result, value)
		}
	}
	return result
}

func hasVowel(word string) bool {
	var vowels = "aeiouAEIOU"
	for _, value := range word {
		if strings.Contains(vowels, string(value)) {
			return true
		}
	}
	return false
}

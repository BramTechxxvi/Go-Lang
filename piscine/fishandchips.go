package piscine

func FishAndChips(n int) string {
	result := ""
	if n%2 == 0 && n%3 == 0 {
		result = "fish and chips"
	} else if n%3 == 0 {
		result = "chips"
	} else if n%2 == 0 {
		result = "fish"
	} else if n < 0 {
		result = "error: number is negative"
	} else if n%2 != 0 && n%3 != 0 {
		result = "error: no divisible"
	}
	return result
}

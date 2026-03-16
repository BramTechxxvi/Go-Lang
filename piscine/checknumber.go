package piscine

func CheckNumber(s string) bool {
	runes := []rune(s)
	for index := 0; index < len(runes); index++ {
		if runes[index] > '0' && runes[index] < '9' {
			return true
		}
	}
	return false
}

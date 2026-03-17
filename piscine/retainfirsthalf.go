package piscine

func RetainFirstHalf(str string) string {
	runes := []rune(str)
	if len(runes) <= 1 {
		return str
	}
	result := ""
	for index := 0; index < len(runes)/2; index++ {
		result += string(runes[index])
	}
	return result
}

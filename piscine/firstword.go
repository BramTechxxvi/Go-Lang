package piscine

func FirstWord(str string) string {
	runes := []rune(str)
	result := ""
	for index := 0; index < len(runes); index++ {
		if runes[index] == ' ' {
			break
		}
		result += string(runes[index])
	}
	return result + "\n"
}

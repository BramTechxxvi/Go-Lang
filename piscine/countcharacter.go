package piscine

func CountChar(s string, c rune) int {
	runes := []rune(s)
	count := 0
	if s == "" {
		return 0
	}

	for index := 0; index < len(runes); index++ {
		if runes[index] == c {
			count++
		}
	}
	return count
}

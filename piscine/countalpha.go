package piscine

func CountAlpha(s string) int {
	runes := []rune(s)
	count := 0

	for index := 0; index < len(runes); index++ {
		if (runes[index] > 'a' && runes[index] < 'z') ||
			(runes[index] > 'A' && runes[index] < 'Z') {
			count++
		}
	}
	return count
}

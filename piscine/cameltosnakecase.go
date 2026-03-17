package piscine

func CamelToSnakeCase(str string) string {
	runes := []rune(str)
	if len(runes) < 1 {
		return str
	}
	for index := 0; index < len(runes)-1; index++ {
		if runes[index] >= 'A' && runes[index] <= 'Z' &&
			runes[index+1] >= 'A' && runes[index+1] <= 'Z' {
			return str
		}
	}

	result := ""
	for index := 0; index < len(runes); index++ {
		if runes[index] >= 'A' && runes[index] <= 'Z' {
			if index != 0 {
				result += "_"
			}
			result += string(runes[index] + 32)
		} else {
			result += string(runes[index])
		}
	}
	return result
}

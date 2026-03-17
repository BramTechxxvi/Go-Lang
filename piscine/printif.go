package piscine

func PrintIf(str string) string {
	if len(str) >= 3 || str == "" {
		return "G" + "\n"
	} else {
		return "Invalid Input" + "\n"
	}
}

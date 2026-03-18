package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 10 {
		Error()
		return
	}
	board := [9][9]int{}

	for row, args := range os.Args[1:] {
		if len(args) != 9 {
			Error()
			return
		}
		for col, char := range args {
			if char >= '1' && char <= '9' {
				num := int(char - '0')
				if !isValid(board, row, col, num) {
					Error()
					return
				}
				board[row][col] = num
			} else if char != '.' {
				Error()
				return
			}
		}
	}

	if solve(&board) {
		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {

				z01.PrintRune(rune(board[row][col] + '0'))
				if col != 8 {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	} else {
		Error()
	}

}

func solve(board *[9][9]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {

			if board[row][col] == 0 {
				for num := 1; num <= 9; num++ {

					if isValid(*board, row, col, num) {
						board[row][col] = num
						if solve(board) {
							return true
						}
						board[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [9][9]int, row, col, num int) bool {
	for index := 0; index < 9; index++ {
		if board[row][index] == num || board[index][col] == num ||
			board[row/3*3+index/3][col/3*3+index%3] == num {
			return false
		}
	}
	return true
}

func Error() {
	for _, r := range "Error\n" {
		z01.PrintRune(r)
	}
}

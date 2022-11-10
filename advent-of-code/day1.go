package adventofcode

import (
	"fmt"
	"strings"
)

type Day1 struct {
	input *string
}

func (day1 *Day1) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day1.part1(), day1.part2()

}

func (day1 *Day1) part1() PuzzleAnswer {
	var ans PuzzleAnswer
	ans.part = 1
	ans.answer = day1.quickParseParens()
	return ans
}

func (day1 *Day1) part2() PuzzleAnswer {
	var ans PuzzleAnswer
	ans.part = 2
	ans.answer = day1.countParens()
	return ans
}

func (day1 *Day1) quickParseParens() string {
	leftParens := strings.Count(*day1.input, "(")
	rightParens := strings.Count(*day1.input, ")")

	return fmt.Sprint(leftParens - rightParens)
}

func (day1 *Day1) countParens() string {
	counter := 0
	for i, char := range *day1.input {
		if char == '(' {
			counter += 1
		}

		if char == ')' {
			counter -= 1
		}

		if counter < 0 {
			return fmt.Sprint(i + 1)
		}
	}
	return ""
}

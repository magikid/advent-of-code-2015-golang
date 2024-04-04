package adventofcode

import (
	"fmt"
	"strings"
)

const (
	niceString    = "nice"
	naughtyString = "naughty"
)

type Day5 struct {
	input *string
}

func (day5 *Day5) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day5.part1(), day5.part2()
}

func (d *Day5) part1() PuzzleAnswer {
	pa := PuzzleAnswer{}
	pa.part = 1
	pa.answer = countNiceStrings(strings.Split(*d.input, "\n"), p1rater)
	return pa
}

func (d *Day5) part2() PuzzleAnswer {
	pa := PuzzleAnswer{}
	pa.part = 2
	pa.answer = countNiceStrings(strings.Split(*d.input, "\n"), p2rater)
	return pa
}

func p1rater(s string) bool {
	return hasAtLeastThreeVowels(s) && hasDoubleLetter(s) && !hasBadStrings(s)
}

func p2rater(s string) bool {
	return hasDoubleDouble(s) && hasRepeatingLetter(s)
}

func countNiceStrings(input []string, rater func(s string) bool) string {
	niceStrings := 0

	for _, s := range input {
		n := niceometer{input: &s, rater: rater}
		if n.rate() == niceString {
			niceStrings++
		}
	}
	return fmt.Sprint(niceStrings) + " nice strings"
}

type niceometer struct {
	input *string
	rater func(s string) bool
}

func (n *niceometer) rate() string {
	if n.rater(*n.input) {
		return niceString
	}

	return naughtyString
}

func hasAtLeastThreeVowels(s string) bool {
	vowels := 0
	for _, r := range s {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}
	}
	return vowels >= 3
}

func hasDoubleLetter(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if (s)[i] == (s)[i+1] {
			return true
		}
	}
	return false
}

func hasBadStrings(s string) bool {
	badStrings := []string{"ab", "cd", "pq", "xy"}
	for _, badString := range badStrings {
		if strings.Contains(s, badString) {
			return true
		}
	}
	return false
}

func hasDoubleDouble(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if strings.Count(s, s[i:i+2]) > 1 {
			return true
		}
	}
	return false
}

func hasRepeatingLetter(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

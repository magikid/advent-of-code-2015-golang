package adventofcode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day2 struct {
	input *string
}

type present struct {
	dimensions string
}

func (present *present) neededWrappingPaper() int64 {
	length, width, height := present.parseDimensions()

	return 2*length*width + 2*width*height + 2*height*length + present.mins(length*width, width*height, height*length)
}

func (p *present) neededRibbon() int64 {
	length, width, height := p.parseDimensions()

	bowLength := length * height * width
	wrappingLength := p.mins((2*length + 2*height), (2*length + 2*width), (2*height + 2*width))

	return bowLength + wrappingLength
}

func (day2 *Day2) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day2.part1(), day2.part2()

}

func (day2 *Day2) part1() PuzzleAnswer {
	var ans PuzzleAnswer
	ans.part = 1
	neededWrappingPaper := int64(0)
	for _, line := range strings.Split(*day2.input, "\n") {
		var present present
		present.dimensions = line
		neededWrappingPaper += present.neededWrappingPaper()
	}

	ans.answer = fmt.Sprint(neededWrappingPaper)
	return ans
}

func (day2 *Day2) part2() PuzzleAnswer {
	var ans PuzzleAnswer
	ans.part = 2
	neededRibbon := int64(0)
	for _, line := range strings.Split(*day2.input, "\n") {
		var present present
		present.dimensions = line
		neededRibbon += present.neededRibbon()
	}

	ans.answer = fmt.Sprint(neededRibbon)
	return ans
}

func (present *present) parseInt(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 0)

	return i
}

func (present *present) mins(value int64, values ...int64) int64 {
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

func (p *present) parseDimensions() (int64, int64, int64) {
	pattern := regexp.MustCompile(`(\d+)x(\d+)x(\d+)`)
	matches := pattern.FindStringSubmatch(p.dimensions)

	if len(matches) == 0 {
		return 0, 0, 0
	}

	length := p.parseInt(matches[1])
	width := p.parseInt(matches[2])
	height := p.parseInt(matches[3])

	return length, width, height
}

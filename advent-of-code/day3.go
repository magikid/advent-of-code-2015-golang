package adventofcode

import (
	"fmt"
)

type Day3 struct {
	input *string
}

type santa struct {
	instructions string
}

func (day3 *Day3) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day3.part1(), day3.part2()
}

func (d *Day3) part1() PuzzleAnswer {
	pa := PuzzleAnswer{}
	pa.part = 1
	directions := santa{instructions: *d.input}
	pa.answer = fmt.Sprint(directions.countDeliveries())
	return pa
}

func (d *Day3) part2() PuzzleAnswer {
	pa := PuzzleAnswer{}
	pa.part = 2
	currentTurn := "santa"
	santaInstructions := ""
	roboSantaInstructions := ""
	for _, instruction := range *d.input {
		switch currentTurn {
		case "santa":
			santaInstructions += string(instruction)
			currentTurn = "roboSanta"
		case "roboSanta":
			roboSantaInstructions += string(instruction)
			currentTurn = "santa"
		default:
			pa.err = fmt.Errorf("an instruction for neither santa nor robosanta came in")
			return pa
		}
	}

	s := santa{instructions: santaInstructions}
	rs := santa{instructions: roboSantaInstructions}
	pa.answer = fmt.Sprint(s.countDeliveries() + rs.countDeliveries() - 1)

	return pa
}

func (s *santa) countDeliveries() int {
	currentX := 500
	currentY := 500
	deltaX := 0
	deltaY := 0
	var houses [1000][1000]int

	houses[currentX][currentY] += 1
	for _, direction := range s.instructions {
		deltaX, deltaY = s.nextPosition(direction)
		currentX += deltaX
		currentY += deltaY
		houses[currentX][currentY] += 1
	}

	housesWithAPresent := 0
	for _, row := range houses {
		for _, col := range row {
			if col > 0 {
				housesWithAPresent += 1
			}
		}
	}

	return housesWithAPresent
}

func (s *santa) nextPosition(direction rune) (int, int) {
	deltaX := 0
	deltaY := 0

	switch direction {
	case '>':
		deltaX = 1
	case '<':
		deltaX = -1
	case '^':
		deltaY = 1
	case 'v':
		deltaY = -1
	}

	return deltaX, deltaY
}

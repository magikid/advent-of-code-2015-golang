package adventofcode

import (
	"flag"
	"fmt"
)

const AllDays = 0
const AllParts = 0

func CLI(args []string) int {
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}

	fmt.Println(app)
	return 0
}

type appEnv struct {
	day             int
	part            int
	puzzleInputPath string
}

func (app appEnv) String() string {
	return fmt.Sprintf("day: %v, part: %v, puzzle input path: %v", app.day, app.part, app.puzzleInputPath)
}

func (app *appEnv) fromArgs(args []string) error {
	fl := flag.NewFlagSet("advent-of-code", flag.ContinueOnError)
	fl.IntVar(&app.day, "day", AllDays, "Which day should be run?")
	fl.IntVar(&app.part, "part", AllParts, "Which part of the given day should be run?")
	fl.StringVar(&app.puzzleInputPath, "path", "inputs/d1p1.txt", "The path to the puzzle input")

	if err := fl.Parse(args); err != nil {
		fl.Usage()
		return err
	}

	return nil
}

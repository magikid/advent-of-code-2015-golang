package adventofcode

import (
	"flag"
	"fmt"
)

const AllDays = 0
const AllParts = 0

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
	fl.StringVar(&app.puzzleInputPath, "path", "inputs", "The path to the puzzle inputs")

	if err := fl.Parse(args); err != nil {
		fl.Usage()
		return err
	}

	return nil
}

func (app *appEnv) run() error {
	// Check that the file path exists and is readable

	// Check that the given day and part exist and aren't empty
	// Build a list of existing and readable inputs
	// For each input, run its runner

	return nil
}

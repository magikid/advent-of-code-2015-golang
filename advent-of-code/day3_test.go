package adventofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtLeastOnePresent(t *testing.T) {
	s := santa{instructions: ">"}
	assert.Equal(t, 2, s.countDeliveries())

	s = santa{instructions: "^>v<"}
	assert.Equal(t, 4, s.countDeliveries())

	s = santa{instructions: "^v^v^v^v^v"}
	assert.Equal(t, 2, s.countDeliveries())
}

func TestRoboSanta(t *testing.T) {
	input := "^v"
	d := Day3{input: &input}
	assert.Equal(t, "3", d.part2().answer)

	input = "^>v<"
	d = Day3{input: &input}
	assert.Equal(t, "3", d.part2().answer)

	input = "^v^v^v^v^v"
	d = Day3{input: &input}
	assert.Equal(t, "11", d.part2().answer)
}

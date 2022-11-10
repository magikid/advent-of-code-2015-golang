package adventofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMins(t *testing.T) {
	present := present{dimensions: "2x3x4"}
	assert.Equal(t, int64(6), present.mins(6, 12, 8))
}

func TestWrappingCalculator(t *testing.T) {
	p := present{dimensions: "2x3x4"}
	assert.Equal(t, int64(58), p.neededWrappingPaper())

	p = present{dimensions: "1x1x10"}
	assert.Equal(t, int64(43), p.neededWrappingPaper())
}

func TestRibbonCalculator(t *testing.T) {
	p := present{dimensions: "2x3x4"}
	assert.Equal(t, int64(34), p.neededRibbon())

	p = present{dimensions: "1x1x10"}
	assert.Equal(t, int64(14), p.neededRibbon())
}

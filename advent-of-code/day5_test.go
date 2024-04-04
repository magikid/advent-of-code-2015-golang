package adventofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	input := "ugknbfddgicrmopn"
	n := niceometer{input: &input, rater: p1rater}
	assert.Equal(t, niceString, n.rate())
}

func Test2(t *testing.T) {
	input := "aaa"
	n := niceometer{input: &input, rater: p1rater}
	assert.Equal(t, niceString, n.rate())
}

func Test3(t *testing.T) {
	input := "jchzalrnumimnmhp"
	n := niceometer{input: &input, rater: p1rater}
	assert.Equal(t, naughtyString, n.rate())
}

func Test4(t *testing.T) {
	input := "haegwjzuvuyypxyu"
	n := niceometer{input: &input, rater: p1rater}
	assert.Equal(t, naughtyString, n.rate())
}

func Test5(t *testing.T) {
	input := "dvszwmarrgswjxmb"
	n := niceometer{input: &input, rater: p1rater}
	assert.Equal(t, naughtyString, n.rate())
}

func Test6(t *testing.T) {
	input := "qjhvhtzxzqqjkmpb"
	n := niceometer{input: &input, rater: p2rater}
	assert.Equal(t, niceString, n.rate())
}

func Test7(t *testing.T) {
	input := "xxyxx"
	n := niceometer{input: &input, rater: p2rater}
	assert.Equal(t, niceString, n.rate())
}

func Test8(t *testing.T) {
	input := "uurcxstgmygtbstg"
	n := niceometer{input: &input, rater: p2rater}
	assert.Equal(t, naughtyString, n.rate())
}

func Test9(t *testing.T) {
	input := "ieodomkazucvgmuy"
	n := niceometer{input: &input, rater: p2rater}
	assert.Equal(t, naughtyString, n.rate())
}

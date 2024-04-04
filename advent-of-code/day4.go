package adventofcode

import (
	"crypto/md5"
	"fmt"
	"strings"
)

type Day4 struct {
	input *string
}

func (day4 *Day4) parts() (PuzzleAnswer, PuzzleAnswer) {
	return day4.part1(), day4.part2()
}

func (d *Day4) part1() PuzzleAnswer {
	pa := PuzzleAnswer{}
	pa.part = 1
	m := miner{secret: *d.input}
	hash, err := m.mineHash("00000")
	if err != nil {
		pa.err = err
		return pa
	}
	pa.answer = hash
	return pa
}

func (d *Day4) part2() PuzzleAnswer {
	pa := PuzzleAnswer{}
	pa.part = 2
	m := miner{secret: *d.input}
	hash, err := m.mineHash("000000")
	if err != nil {
		pa.err = err
		return pa
	}
	pa.answer = hash
	return pa
}

type miner struct {
	secret string
}

func (m *miner) mineHash(prefix string) (string, error) {
	var b strings.Builder
	for i := uint(1); i < ^uint(0); i++ {
		b.WriteString(fmt.Sprintf("%s%d", m.secret, i))
		if strings.HasPrefix(m.hash(b.String()), prefix) {
			return fmt.Sprintf("%d", i), nil
		}
		b.Reset()
	}

	return "", fmt.Errorf("no hash found")
}

func (m *miner) hash(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

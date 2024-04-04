package adventofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashing(t *testing.T) {
	m := miner{secret: "abcdef"}
	hash, err := m.mineHash()
	assert.Nil(t, err)
	assert.Equal(t, "609043", hash)
}

func TestHashing2(t *testing.T) {
	m := miner{secret: "pqrstuv"}
	hash, err := m.mineHash()
	assert.Nil(t, err)
	assert.Equal(t, "1048970", hash)
}

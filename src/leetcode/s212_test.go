package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test212(t *testing.T) {
	assert := assert.New(t)
	board := [][]byte{
		[]byte("oaan"),
		[]byte("etae"),
		[]byte("ihkr"),
		[]byte("ihkr"),
	}

	words := []string{"oath", "pea", "eat", "rain"}
	assert.Equal(FindWords(words, board), []string{"eat", "oath"}, "should return []string{\"oath\",\"eat\"}")
}

func Test212_2(t *testing.T) {
	assert := assert.New(t)
	board := [][]byte{
		[]byte("a"),
	}

	words := []string{"a", "a"}
	assert.Equal(FindWords(words, board), []string{"a"}, "should return a")
}

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

	nws := NewWordSearch2(board)
	words := []string{"oath", "pea", "eat", "rain"}
	assert.Equal(nws.FindWords(words), []string{"eat", "oath"}, "should return []string{\"oath\",\"eat\"}")
}

func Test212_2(t *testing.T) {
	assert := assert.New(t)
	board := [][]byte{
		[]byte("a"),
	}

	nws := NewWordSearch2(board)
	words := []string{"a", "a"}
	assert.Equal(nws.FindWords(words), []string{"a"}, "should return a")
}

package backtrack_test

import (
	"testing"

	"github.com/h4ckm03d/coding-practice/backtrack"
	"github.com/stretchr/testify/assert"
)

func TestCaseA(t *testing.T) {
	assert := assert.New(t)
	board := [][]byte{
		[]byte("CAA"),
		[]byte("AAA"),
		[]byte("BCD"),
	}

	nws := backtrack.NewWordSearch(board)

	assert.Equal(nws.Exist("AAB"), true, "should return true")
}

func TestCaseB(t *testing.T) {
	assert := assert.New(t)
	board := [][]byte{
		[]byte("oaan"),
		[]byte("etae"),
		[]byte("ihkr"),
		[]byte("ihkr"),
	}

	nws := backtrack.NewWordSearch(board)
	words := []string{"oath", "pea", "eat", "rain"}
	assert.Equal(nws.FindWords(words), []string{"eat", "oath"}, "should return []string{\"oath\",\"eat\"}")
}

func TestCaseC(t *testing.T) {
	assert := assert.New(t)
	board := [][]byte{
		[]byte("a"),
	}

	nws := backtrack.NewWordSearch(board)
	words := []string{"a", "a"}
	assert.Equal(nws.FindWords(words), []string{"a"}, "should return a")
}

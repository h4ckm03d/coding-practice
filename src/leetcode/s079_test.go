package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test079(t *testing.T) {
	assert := assert.New(t)
	board := [][]byte{
		[]byte("CAA"),
		[]byte("AAA"),
		[]byte("BCD"),
	}

	nws := NewWordSearch(board)

	assert.Equal(nws.Exist("AAB"), true, "should return true")
	assert.Equal(nws.Exist("AABD"), false, "should return false")
}
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
	assert.Equal(FindWord("AAB", board), true, "should return true")
	assert.Equal(FindWord("AABD", board), false, "should return false")
}

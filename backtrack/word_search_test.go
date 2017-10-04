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
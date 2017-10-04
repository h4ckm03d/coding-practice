package backtrack_test

import (
	"testing"
	"github.com/h4ckm03d/coding-practice/backtrack"
)

func TestCaseA(t *testing.T) {
	board := [][]byte{
		[]byte("CAA"),
		[]byte("AAA"),
		[]byte("BCD"),
	}

	nws := backtrack.NewWordSearch(board)
	
	if nws.Exist("AAB")==true{
		t.Log("Success")
	}else{
		t.Error("Must return true")
	}
}
package dfs_test

import (
	"testing"

	"github.com/h4ckm03d/coding-practice/dfs"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert := assert.New(t)
	matrix := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	assert.Equal(4, dfs.LongestIncreasingPath(matrix), " should return 4")
}

func BenchmarkLip(b *testing.B) {
	matrix := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	for n := 0; n < b.N; n++ {
		dfs.LongestIncreasingPath(matrix);
	}
}

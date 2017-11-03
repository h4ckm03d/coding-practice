package leetcode

// Given an integer matrix, find the length of the longest increasing path.
// From each cell, you can either move to four directions: left, right, up or down. You may NOT move diagonally or move outside of the boundary (i.e. wrap-around is not allowed).
// Example 1:
// nums = [
//   [9,9,4],
//   [6,6,8],
//   [2,1,1]
// ]
// Return 4
// The longest increasing path is [1, 2, 6, 9].
// Example 2:
// nums = [
//   [3,4,5],
//   [3,2,6],
//   [2,2,1]
// ]
// Return 4
// The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

//LongestIncreasingPath find longest increasing path
func LongestIncreasingPath(matrix [][]int) int {
	m := len(matrix)

	if m == 0 {
		return 0
	}
	
	n := len(matrix[0])
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
	}
	max := 1
	for i := 0; i < m; i++{
		for j := 0; j < n; j++ {
			len := dfs(matrix, i, j, m, n, cache)
			max = Max(max, len)
		}
	}
	return max
}

func dfs(matrix [][]int, i, j, m, n int, cache [][]int) int {
	
	if cache[i][j] != 0 {
		return cache[i][j]
	}
	max := 1
	for _, dir := range dirs {
		x := i + dir[0]
		y := j + dir[1]
		if x < 0 || x >= m || y < 0 || y >= n || matrix[x][y] <= matrix[i][j] {
			continue
		}
		len := 1 + dfs(matrix, x, y, m, n, cache)
		max = Max(max, len)
	}
	cache[i][j] = max
	return max
}

// Max to get bigger value
func Max(a, b int) int {
	if a <= b {
		return b
	}

	return a
}

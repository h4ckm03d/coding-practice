package leetcode

// WordSearch is solution for word search problem
// Given a 2D board and a word, find if the word exists in the grid.
// The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.
// For example,
// Given board =
// [
//   ['A','B','C','E'],
//   ['S','F','C','S'],
//   ['A','D','E','E']
// ]
// word = "ABCCED", -> returns true,
// word = "SEE", -> returns true,
// word = "ABCB", -> returns false.
type WordSearch struct {
	board [][]byte
}

// Exist is for checking existing word within boards
func (w *WordSearch) Exist(word string) bool {
	for i, h := range w.board {
		for j := range h {
			if w.dfs(w.board, word, i, j) == true {
				return true
			}
		}
	}
	return false
}

// dfs deep first search
func (w *WordSearch) dfs(board [][]byte, word string, i, j int) bool {
	//fmt.Println(board)
	if len(word) == 0 {
		return true
	}

	if i < 0 ||
		i >= len(board) ||
		j < 0 ||
		j >= len(board[0]) ||
		board[i][j] != word[0] {
		return false
	}
	//fmt.Println(board, word, i, j, len(board),len(board[0]), board[i][j]!=word[0], word[0], word[1:], board[i][j]  )
	tmp := board[i][j]
	board[i][j] = '#'
	result := w.dfs(board, word[1:], i+1, j) ||
		w.dfs(board, word[1:], i-1, j) ||
		w.dfs(board, word[1:], i, j+1) ||
		w.dfs(board, word[1:], i, j-1)
	board[i][j] = tmp
	return result
}

// NewWordSearch for create new instance wordsearch
func NewWordSearch(board [][]byte) *WordSearch {
	return &WordSearch{
		board: board,
	}
}
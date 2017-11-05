package leetcode

import (
	"sort"
)

// WordSearch2 is solution for word search problem
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
type WordSearch2 struct {
	board [][]byte
}

// WordExist is for checking existing word within boards
func WordExist(word string, board [][]byte) bool {
	for i, h := range board {
		for j := range h {
			if dfsWS(board, word, i, j) == true {
				return true
			}
		}
	}
	return false
}

// dfsWS deep first search
func dfsWS(board [][]byte, word string, i, j int) bool {
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
	result := dfsWS(board, word[1:], i+1, j) ||
		dfsWS(board, word[1:], i-1, j) ||
		dfsWS(board, word[1:], i, j+1) ||
		dfsWS(board, word[1:], i, j-1)
	board[i][j] = tmp
	return result
}

// FindWords to find words from boards
// Given a 2D board and a list of words from the dictionary, find all words in the board.
// Each word must be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.
// For example,
// Given words = ["oath","pea","eat","rain"] and board =
// [
//   ['o','a','a','n'],
//   ['e','t','a','e'],
//   ['i','h','k','r'],
//   ['i','f','l','v']
// ]
// Return ["eat","oath"].
// Note:
// You may assume that all inputs are consist of lowercase letters a-z.
// You would need to optimize your backtracking to pass the larger test. Could you stop backtracking earlier?
// If the current candidate does not exist in all words' prefix, you could stop backtracking immediately. What kind of data structure could answer such query efficiently? Does a hash table work? Why or why not? How about a Trie? If you would like to learn how to implement a basic trie, please work on this problem: Implement Trie (Prefix Tree) first.
func FindWords(words []string, board [][]byte) []string {
	words = removeDuplicates(words)
	sort.Strings(words)
	var results = []string{}
	for _, word := range words {
		if WordExist(word, board) {
			results = append(results, word)
		}
	}

	return results
}

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

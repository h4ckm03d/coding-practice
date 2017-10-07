package string

// LengthOfLongestSubstring get total number substring
func LengthOfLongestSubstring(s string) int {
	n, ans := len(s), 0
	m := make(map[byte]int)
	i, j := 0, 0
	for ; j < n; j++ {
		if val, ok := m[s[j]]; ok {
			i = Max(val, i)
		}
		ans = Max(ans, j-i+1)
		m[s[j]] = j + 1
	}
	return ans
}

// Max to get bigger value
func Max(a, b int) int {
	if a <= b {
		return b
	}

	return a
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test003(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(3, LengthOfLongestSubstring("abcabcbb"), "It should return 3")

	assert.Equal(6, Max(6, 4), "")
}

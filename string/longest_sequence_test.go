//abcabcbb
package string_test

import (
	"testing"

	"github.com/h4ckm03d/coding-practice/string"
	"github.com/stretchr/testify/assert"
)

func TestLSC(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(3, string.LengthOfLongestSubstring("abcabcbb"), "It should return 3")

	assert.Equal(6, string.Max(6, 4), "")
}

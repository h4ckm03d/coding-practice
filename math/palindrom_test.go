package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegerPalindrom(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(true, IsPalindrome(1221), "")
	assert.Equal(false, IsPalindrome(-1), "")
	assert.Equal(true, IsPalindrome(11), "")
	assert.Equal(true, IsPalindrome2(11), "")
	assert.Equal(true, IsPalindrome2(1221), "")
	assert.Equal(false, IsPalindrome2(-1), "")
}

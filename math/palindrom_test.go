package math_test

import (
	"testing"

	"github.com/h4ckm03d/coding-practice/math"
	"github.com/stretchr/testify/assert"
)

func TestIntegerPalindrom(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(true, math.IsPalindrome(1221), "")
	assert.Equal(false, math.IsPalindrome(-1), "")
	assert.Equal(true, math.IsPalindrome(11), "")
	assert.Equal(true, math.IsPalindrome2(11), "")
	assert.Equal(true, math.IsPalindrome2(1221), "")
	assert.Equal(false, math.IsPalindrome2(-1), "")
}

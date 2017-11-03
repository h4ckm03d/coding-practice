package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLru2(t *testing.T) {
	assert := assert.New(t)

	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	assert.Equal(c.Get(1), 1, "should return 1")
	c.Put(3, 3)
	c.Put(3, 3)
	assert.Equal(c.Get(2), -1, "should return -1")
	c.Put(4, 4)
	assert.Equal(c.Get(1), -1, "should return -1")
	assert.Equal(c.Get(3), 3, "should return 3")
	assert.Equal(c.Get(4), 4, "should return 4")
}

package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test003(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(5, Solution(1041))
	assert.Equal(0, Solution(15))
	assert.Equal(29, Solution(1073741825 ))
	assert.Equal(5, Solution(1376796946 ))
}

func Benchmark003(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Solution(n)
	}
}

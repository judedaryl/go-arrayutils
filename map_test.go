package arrayutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	var arr = []int{1, 2, 3, 4}
	result := Map(arr, func(val int) int {
		return val + 1
	})

	// should not modify array length
	assert.Equal(t, 4, len(arr))
	// should be increased by 1
	assert.Equal(t, 2, result[0])
	assert.Equal(t, 3, result[1])
	assert.Equal(t, 4, result[2])
	assert.Equal(t, 5, result[3])
}

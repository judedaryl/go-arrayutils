package arrayutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	var arr = []int{1, 2, 3, 4}
	evenOnly := Filter(arr, func(val int) bool {
		return val%2 == 0
	})

	// should not modify array length
	assert.Equal(t, 2, len(evenOnly))
	// should be increased by 1
	assert.Equal(t, 2, evenOnly[0])
	assert.Equal(t, 4, evenOnly[1])
}

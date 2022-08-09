package arrayutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	var arr = []int{1, 2, 3, 4}
	firstEven := Find(arr, func(val int) bool {
		return val%2 == 0
	})
	// should not modify array length
	assert.Equal(t, 2, *firstEven)
	shouldNil := Find(arr, func(val int) bool {
		return val == 5
	})
	assert.Nil(t, shouldNil)
}

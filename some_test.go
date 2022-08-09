package arrayutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSome(t *testing.T) {
	var arr = []int{1, 2, 3, 4}
	falseResult := Some(arr, func(val int) bool {
		return val > 4
	})

	assert.Equal(t, false, falseResult)

	trueResult := Some(arr, func(val int) bool {
		return val == 4
	})

	assert.Equal(t, true, trueResult)
}

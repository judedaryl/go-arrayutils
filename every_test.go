package arrayutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvery(t *testing.T) {
	var arr = []int{1, 2, 3, 4}
	falseResult := Every(arr, func(val int) bool {
		return val > 4
	})

	assert.Equal(t, false, falseResult)

	stillFalse := Every(arr, func(val int) bool {
		return val == 4
	})

	assert.Equal(t, false, stillFalse)

	shouldTrue := Every(arr, func(val int) bool {
		return val < 5
	})

	assert.Equal(t, true, shouldTrue)
}

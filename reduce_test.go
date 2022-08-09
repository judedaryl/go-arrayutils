package arrayutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	intarr := []int{2, 2, 3, 4}
	sum := Reduce(intarr, 0, func(agg int, val int) int {
		return agg + val
	})
	assert.Equal(t, 11, sum)
}

package arrayutils

/*
Invokes `aggregator` for each element in the array and
reduces the values of an array to a single value (going left-to-right)

Example:
	intarr := []int{1, 2, 3, 4}
	sum := Reduce(intarr, 0, func(agg int, val int) int {
		return agg + val
	})
*/
func Reduce[TIn any, TAgg any](arr []TIn, initialValue TAgg, reducer func(agg TAgg, current TIn) TAgg) TAgg {
	agg := initialValue
	for i := range arr {
		agg = reducer(agg, arr[i])
	}
	return agg
}

package arrayutils

/*
Invokes `predicate` for each element in the array.
checks if every array elements pass a test (provided as a callback function).

Example:
	intarr := []int{1, 2, 3, 4}
	shouldTrue := Every(intarr, func(val int) bool {
		return val < 5
	})
	shouldFalse := Every(intarr, func(val int) bool {
		return val > 1
	})
*/
func Every[TIn any](arr []TIn, predicate func(val TIn) bool) bool {
	result := true
	for i := range arr {
		result = result && predicate(arr[i])
	}
	return result
}

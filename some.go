package arrayutils

/*
Invokes `predicate` for each element in the array.
checks if any array elements pass a test (provided as a callback function).

Example:
	var arr = []int{1, 2, 3, 4}
	falseResult := Some(arr, func(val int) bool {
		return val > 4
	})
*/
func Some[TIn any](arr []TIn, predicate func(val TIn) bool) bool {
	for i := range arr {
		if predicate(arr[i]) {
			return true
		}
	}
	return false
}

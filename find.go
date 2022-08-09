package arrayutils

/*
Invokes `predicate` for each element in the array and returning
the first elements to pass the test

Example:
	intarr := []int{1, 2, 3, 4}
	firstEven := Find(intarr, func(val int) bool {
		return val%2 == 0
	})
*/
func Find[TIn any](arr []TIn, predicate func(val TIn) bool) *TIn {
	for i := range arr {
		if predicate(arr[i]) {
			return &arr[i]
		}
	}
	return nil
}

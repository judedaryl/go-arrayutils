package arrayutils

/*
Invokes `predicate` for each element in the array and removing
any element where predicate returns a false

Example:
	intarr := []int{1, 2, 3, 4}
	onlyTwo := Filter(intarr, func(val int) bool {
		return val == 2
	})
*/
func Filter[TIn any](arr []TIn, predicate func(val TIn) bool) []TIn {
	output := []TIn{}
	for i := range arr {
		if predicate(arr[i]) {
			output = append(output, arr[i])
		}
	}
	return output
}

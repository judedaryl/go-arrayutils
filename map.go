package arrayutils

/*
Invokes `selector` for each element in the array transforming them to a form
dictated by the selector

Example:
	intarr := []int{1, 2, 3, 4}
	stringarr := Map(intarr, func(val int) string {
		return strconv.Itoa(val)
	})
*/
func Map[TIn any, TOut any](arr []TIn, selector func(val TIn) TOut) []TOut {
	output := []TOut{}
	for i := range arr {
		out := selector(arr[i])
		output = append(output, out)
	}
	return output
}

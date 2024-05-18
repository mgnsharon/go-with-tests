package arraysslices

// Sum calculates the sum of a slice of integers
func Sum(nums []int) (sum int) {

	for _, v := range nums {
		sum += v
	}
	return sum
}

// SumAll calculates the sums of multiple slices of integers
func SumAll(nums [][]int) (totals []int) {
	for _, items := range nums {
		totals = append(totals, Sum(items))
	}
	return totals
}

// SumVariadic calculates the sums of multiple slices of integers passed as variadic parameters
func SumVariadic(nums ...[]int) []int {
	totals := make([]int, len(nums))
	for i, v := range nums {
		totals[i] = Sum(v)
	}
	return totals
}

// SumTails calculates the sum of all elements except the first one for multiple slices of integers
func SumTails(nums ...[]int) []int {
	totals := make([]int, len(nums))
	for i, v := range nums {
		if len(v) == 0 {
			totals[i] = 0
		} else {
			totals[i] = Sum(v[1:])
		}
	}
	return totals
}

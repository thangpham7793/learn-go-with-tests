package array_and_slices

func Sum(nums [5]int) int {
	sum := 0

	for _, v := range nums {
		sum += v
	}

	return sum
}

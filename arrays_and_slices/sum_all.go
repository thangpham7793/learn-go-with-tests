package array_and_slices

func SumAll(args ...[]int) []int {
	var res []int

	for _, v := range args {
		s, _ := Sum(v)
		res = append(res, s)
	}

	return res
}

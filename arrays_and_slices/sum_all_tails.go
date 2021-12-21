package array_and_slices

func SumAllTails(args ...[]int) []int {
	var res []int

	for _, v := range args {
		// if v has len 1 then [1:] creates [0]
		s, _ := Sum(v[1:])
		res = append(res, s)
	}

	return res
}

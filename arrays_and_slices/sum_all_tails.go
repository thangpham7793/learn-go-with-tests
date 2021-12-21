package array_and_slices

import "errors"

var EmptySliceError = errors.New("no empty slice allowed")

func SumAllTails(args ...[]int) ([]int, error) {
	var res []int

	for _, v := range args {
		if len(v) == 0 {
			return nil, EmptySliceError
		}

		// if v has len 1 then [1:] creates [0]
		s, _ := Sum(v[1:])
		res = append(res, s)
	}

	return res, nil
}

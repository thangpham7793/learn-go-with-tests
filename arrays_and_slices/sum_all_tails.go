package array_and_slices

import (
	"errors"
)

var EmptySliceError = errors.New("no empty slice allowed")
var NoSliceError = errors.New("must contain at least one slice")

func SumAllTails(containers ...[]int) ([]int, error) {
	var res []int
	if len(containers) == 0 {
		return nil, EmptySliceError
	}

	for _, v := range containers {
		if len(v) == 0 {
			return nil, EmptySliceError
		}

		// if v has len 1 then [1:] creates [0]
		s, _ := Sum(v[1:])
		res = append(res, s)
	}

	return res, nil
}

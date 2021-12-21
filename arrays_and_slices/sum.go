package array_and_slices

import (
	"errors"
)

var InvalidTypeError = errors.New("must be a slice or array of integers")
var EmptyCollectionError = errors.New("collection must have at least 1 element")

func Sum(nums interface{}) (int, error) {
	sum := 0
	switch nums.(type) {
	case [5]int:
		if len(nums.([5]int)) == 0 {
			return 0, EmptyCollectionError
		}
		for _, v := range nums.([5]int) {
			sum += v
		}
	case []int:
		if len(nums.([]int)) == 0 {
			return 0, EmptyCollectionError
		}
		for _, v := range nums.([]int) {
			{
				sum += v
			}
		}
	default:
		return 0, InvalidTypeError
	}

	return sum, nil
}

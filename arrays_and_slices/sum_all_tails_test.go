package array_and_slices

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	t.Run("can handle slices of various length", func(t *testing.T) {
		testCases := []test{
			{
				input: [][]int{{1}, {3, 4}},
				want:  []int{0, 4},
			},
			{
				input: [][]int{{1, 2, 3}, {3, 4, 5}},
				want:  []int{5, 9},
			},
		}

		for _, tc := range testCases {
			args := tc.input.([][]int)
			got, _ := SumAllTails(args...)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %d want %d given %v", got, tc.want, tc.input)
			}
		}
	})

	t.Run("throw error if any slice is empty", func(t *testing.T) {
		if _, err := SumAllTails([][]int{}...); err == nil {
			t.Error("should throw error if any slice is empty")
		}
	})

	t.Run("throw error if container slice is empty", func(t *testing.T) {
		if _, err := SumAllTails([][]int{}...); err == nil {
			t.Error("should throw error if container slice is empty")
		}
	})
}

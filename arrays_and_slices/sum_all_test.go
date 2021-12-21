package array_and_slices

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	testCases := []test{
		{
			input: [][]int{{1, 2}, {3, 4}},
			want:  []int{3, 7},
		},
		{
			input: [][]int{{1, 2}, {3, 4}, {5, 6}},
			want:  []int{3, 7, 11},
		},
	}

	for _, tc := range testCases {
		args := tc.input.([][]int)
		got := SumAll(args...)

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got %d want %d given %v", got, tc.want, tc.input)
		}
	}
}

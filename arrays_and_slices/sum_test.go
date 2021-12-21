package array_and_slices

import (
	"testing"
)

type test struct {
	input interface{}
	want  int
}

func TestSum(t *testing.T) {
	t.Run("can handle both slice & array", func(t *testing.T) {
		testCases := []test{
			{
				input: [5]int{1, 2, 3, 4, 5},
				want:  15,
			},
			{
				input: []int{1, 2, 3, 4, 5, 6, 7},
				want:  28,
			},
		}

		for _, tc := range testCases {
			got, _ := Sum(tc.input)
			if got != tc.want {
				t.Errorf("got %d want %d given, %v", got, tc.want, tc.input)
			}
		}
	})

	t.Run("throw error when its wrong type", func(t *testing.T) {
		numbers := []float32{1, 2, 3, 4, 5, 6, 7}

		if _, err := Sum(numbers); err == nil {
			t.Error("should throw error when ele its not of type int")
		}
	})

	t.Run("throw error when collection (slice) is empty", func(t *testing.T) {
		if _, err := Sum([]int{}); err == nil {
			t.Error("should throw error when collection is empty")
		}
	})
}

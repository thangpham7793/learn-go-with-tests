package array_and_slices

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got, _ := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}
		got, _ := Sum(numbers)

		want := 28

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
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

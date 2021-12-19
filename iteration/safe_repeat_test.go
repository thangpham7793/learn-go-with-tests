package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Should accept a number of time to repeat", func(t *testing.T) {
		want := "ababab"
		got, _ := SafeRepeat("ab", 3)
		if want != got {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	t.Run("Should throw error if number is 0 or negative", func(t *testing.T) {
		_, err := SafeRepeat("ab", 0)
		if err.Error() != "repeat: invalid repeat time: 0" {
			t.Error("Should throw invalid number of time error")
		}
	})
}

func BenchmarkSafeRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SafeRepeat("benchmarkinggg", 1)
	}
}

func ExampleSafeRepeat() {
	res, _ := SafeRepeat("ab", 2)
	fmt.Println(res)
	// Output: abab
}

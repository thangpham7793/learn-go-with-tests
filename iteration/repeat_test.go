package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Should accept a number of time to repeat", func(t *testing.T) {
		want := "ababab"
		got, _ := Repeat("ab", 3)
		if want != got {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})
}

func ExampleRepeat() {
	res, _ := Repeat("ab", 2)
	fmt.Println(res)
	// Output: abab
}

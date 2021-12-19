package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	want := "ababab"
	got, _ := Repeat("ab")
	if want != got {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func ExampleRepeat() {
	res, _ := Repeat("ab")
	fmt.Println(res)
	// Output: ababab
}

package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 4
	if sum := Add(2, 2); sum != 4 {
		t.Errorf("Expected %q, got %v", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(3, 4)
	fmt.Println(sum)
	// Output: 7
}

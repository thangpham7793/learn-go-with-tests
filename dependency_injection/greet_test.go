package dependency_injection

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

// Fprintf takes in a writer, whereas Printf uses os.Stdout by default
// using interface allows a different implementation to be used in test
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

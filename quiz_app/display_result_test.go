package quiz_app

import (
	"bytes"
	"testing"
)

func TestDisplayResults(t *testing.T) {
	buffer := &bytes.Buffer{}

	displayResult(buffer, 3, 10)

	want := "You answered 3 out of 10 questions"
	got := buffer.String()

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

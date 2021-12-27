package quiz_app

import (
	"os"
	"testing"
)

func TestQuiz(t *testing.T) {
	questions := []Question{
		{
			prompt: "1+1",
			answer: "2",
		},
		{
			prompt: "1+1",
			answer: "2",
		},
		{
			prompt: "1+1",
			answer: "2",
		},
	}

	t.Run("should handle all correct answers", func(t *testing.T) {
		r, w, _ := os.Pipe()
		w.WriteString("2\n2\n2\n")

		defer func() {
			w.Close()
		}()

		want := 3
		got, _ := quiz(questions, r)
		if got != want {
			t.Errorf("want %d correct answers, got %d", want, got)
		}
	})

	t.Run("should handle some incorrect answers", func(t *testing.T) {
		r, w, _ := os.Pipe()
		w.WriteString("2\n1\n2\n")

		defer func() {
			w.Close()
		}()

		want := 2
		got, _ := quiz(questions, r)
		if got != want {
			t.Errorf("want %d correct answers, got %d", want, got)
		}
	})
}

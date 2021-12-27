package quiz_app

import (
	"os"
	"strings"
	"testing"
	"time"
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

	simulateUserInput := func(t *testing.T, answers []string) *os.File {
		t.Helper()
		reader, writer, _ := os.Pipe()
		writer.WriteString("\r\n" + strings.Join(answers, "\n"))

		defer func() {
			writer.Close()
		}()

		return reader
	}

	t.Run("should handle all correct answers", func(t *testing.T) {
		reader := simulateUserInput(t, []string{"2", "2", "2"})

		want := 3
		got, _ := quiz(questions, reader, time.Second)
		if got != want {
			t.Errorf("want %d correct answers, got %d", want, got)
		}
	})

	t.Run("should handle some incorrect answers", func(t *testing.T) {
		reader := simulateUserInput(t, []string{"2", "1", "2"})

		want := 2
		got, _ := quiz(questions, reader, time.Second)
		if got != want {
			t.Errorf("want %d correct answers, got %d", want, got)
		}
	})

	t.Run("should exit when time allowed is expired", func(t *testing.T) {
		reader := simulateUserInput(t, []string{"2", "2", "2"})

		want := 0
		// should exit immediately
		got, _ := quiz(questions, reader, 0)
		if got != want {
			t.Errorf("want %d correct answers, got %d", want, got)
		}
	})
}

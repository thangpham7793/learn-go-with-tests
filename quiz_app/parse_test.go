package quiz_app

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func makeTempFile(t *testing.T, lines, format string) *os.File {
	t.Helper()

	dir, _ := os.Getwd()
	f, _ := ioutil.TempFile(dir, fmt.Sprintf("quiz-*.%s", format))
	f.WriteString(lines)
	f.Close()

	file, _ := os.Open(f.Name())
	defer os.Remove(f.Name())

	return file
}

func TestParse(t *testing.T) {

	t.Run("can parse properly formatted csv question set", func(t *testing.T) {
		file := makeTempFile(t, "1+1,2\n2+2,3", "csv")

		want := []Question{{prompt: "1+1", answer: "2"}, {prompt: "2+2", answer: "3"}}
		got, _ := parseCsv(file)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %s want %s", got, want)
		}
	})

	t.Run("can parse formatted csv question set with extra space", func(t *testing.T) {
		file := makeTempFile(t, "1+1,  2 ", "csv")

		want := []Question{{prompt: "1+1", answer: "2"}}
		got, _ := parseCsv(file)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %s want %s", got, want)
		}
	})

	t.Run("return an error if no question is given, or a question has no prompt or answer", func(t *testing.T) {
		testCases := []string{"1+3", "1+1,", ",2", ""}
		for _, lines := range testCases {
			file := makeTempFile(t, lines, "csv")
			questions, error := parseCsv(file)

			if error == nil {
				t.Errorf("case %q: should return error, got %s & no error instead", lines, questions)
			}
		}
	})
}

func TestNewQuestionParser(t *testing.T) {
	t.Run("should return csv parser if format is supported", func(t *testing.T) {
		file := makeTempFile(t, "foo", "csv")
		parser, _ := NewQuestionParser(file)

		if parser == nil {
			t.Error("should return a non-nil parser")
		}
	})

	t.Run("should return nil parser & error if format is not supported", func(t *testing.T) {
		file := makeTempFile(t, "foo", "toml")
		_, error := NewQuestionParser(file)

		if error == nil {
			t.Errorf("should return %q error", UnsupportedFormatError{"toml"})
		}
	})
}

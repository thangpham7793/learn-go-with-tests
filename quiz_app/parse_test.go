package quiz_app

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {

	makeTempFile := func(t *testing.T, lines string) *os.File {
		t.Helper()

		dir, _ := os.Getwd()
		f, _ := ioutil.TempFile(dir, "quiz")
		f.WriteString(lines)
		f.Close()

		file, _ := os.Open(f.Name())
		defer os.Remove(f.Name())

		return file
	}

	t.Run("can parse properly formatted csv question set", func(t *testing.T) {
		file := makeTempFile(t, "1+1,2\n2+2,3")

		want := []Question{{prompt: "1+1", answer: "2"}, {prompt: "2+2", answer: "3"}}
		got, _ := parse(file)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %s want %s", got, want)
		}
	})

	t.Run("can parse formatted csv question set with extra space", func(t *testing.T) {
		file := makeTempFile(t, "1+1,  2 ")

		want := []Question{{prompt: "1+1", answer: "2"}}
		got, _ := parse(file)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %s want %s", got, want)
		}
	})

	t.Run("return an error if no question is given, or a question has no prompt or answer", func(t *testing.T) {
		testCases := []string{"1+3", "1+1,", ",2", ""}
		for _, lines := range testCases {
			file := makeTempFile(t, lines)
			questions, error := parse(file)

			if error == nil {
				t.Errorf("case %q: should return error, got %s & no error instead", lines, questions)
			}
		}
	})
}

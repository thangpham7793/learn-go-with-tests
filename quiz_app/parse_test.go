package quiz_app

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {

	dir, _ := os.Getwd()
	f, _ := ioutil.TempFile(dir, "quiz")
	f.WriteString("1+1,2")

	defer os.Remove(f.Name())

	want := []Question{{prompt: "1+1", answer: "2"}}
	got, _ := parse(f)

	if reflect.DeepEqual(got, want) {
		t.Errorf("Got %s want %s", got, want)
	}

}

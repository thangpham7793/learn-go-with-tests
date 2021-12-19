package hello_world

import "testing"

func TestHello(t *testing.T) {

	testSuites := map[string]string{
		"Hello":   "Thang",
		"Bonjour": "Tommy",
	}

	for greeting, name := range testSuites {
		got := Hello(greeting, name)
		want := greeting + " " + name + "!"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

}

func TestMakeHello(t *testing.T) {

	testSuites := map[string]string{
		"Hello":   "Thang",
		"Bonjour": "Tommy",
	}

	for greeting, name := range testSuites {
		Greet := HelloWith(greeting)
		got := Greet(name)
		want := greeting + " " + name + "!"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

}

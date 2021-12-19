package hello_world

import "testing"

func TestHelloSuite(t *testing.T) {
	t.Run("Say hello to a given name", func(t *testing.T) {
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
	},
	)

	t.Run("Use a different greeting when supplied", func(t *testing.T) {

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
	})

	t.Run("Use 'World' when empty name is given", func(t *testing.T) {

		testSuites := map[string]string{
			"Hello": "",
		}

		for greeting, name := range testSuites {
			got := Hello(greeting, name)
			want := "Hello World!"
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		}
	})

	t.Run("Use 'Hello' when empty greeting is given", func(t *testing.T) {

		testSuites := map[string]string{
			"": "Thang",
		}

		for greeting, name := range testSuites {
			got := Hello(greeting, name)
			want := "Hello Thang!"
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		}
	})
}

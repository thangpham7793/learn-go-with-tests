package hello_world

import "testing"

func TestHelloSuite(t *testing.T) {
	assert := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Say hello to a given name", func(t *testing.T) {
		testSuites := map[string][]string{
			"Hello Thang!":   {"Hello", "Thang"},
			"Bonjour Thang!": {"Bonjour", "Thang"},
		}

		for want, args := range testSuites {
			got := Hello(args[0], args[1])
			assert(t, got, want)
		}
	},
	)

	t.Run("Use a different greeting when supplied", func(t *testing.T) {
		testSuites := map[string][]string{
			"Hello Thang!":   {"Hello", "Thang"},
			"Bonjour Thang!": {"Bonjour", "Thang"},
		}

		for want, args := range testSuites {
			Greet := HelloWith(args[0])
			got := Greet(args[1])
			assert(t, got, want)
		}
	})

	t.Run("Use 'World' when empty name is given", func(t *testing.T) {
		testSuites := map[string][]string{
			"Hello World!": {"Hello", ""},
		}

		for want, args := range testSuites {
			got := Hello(args[0], args[1])
			assert(t, got, want)
		}
	})

	t.Run("Use 'Hello' when empty greeting is given", func(t *testing.T) {
		testSuites := map[string][]string{
			"Hello Thang!": {"", "Thang"},
		}

		for want, args := range testSuites {
			got := Hello(args[0], args[1])
			assert(t, got, want)
		}
	})
}

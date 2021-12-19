package hello_world

import "testing"

func TestHelloSuite(t *testing.T) {
	assert := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Say hello to a given name based on the supplied language", func(t *testing.T) {
		testSuites := map[string][]string{
			"Hello Thang!":   {"English", "Thang"},
			"Bonjour Thang!": {"French", "Thang"},
		}

		for want, args := range testSuites {
			got, _ := Hello(args[0], args[1])
			assert(t, got, want)
		}
	},
	)

	t.Run("Can pre-apply with a given language", func(t *testing.T) {
		testSuites := map[string][]string{
			"Hello Thang!":   {"English", "Thang"},
			"Bonjour Thang!": {"French", "Thang"},
		}

		for want, args := range testSuites {
			Greet := HelloIn(args[0])
			got, _ := Greet(args[1])
			assert(t, got, want)
		}
	})

	t.Run("Use 'World' when empty name is given", func(t *testing.T) {
		testSuites := map[string][]string{
			"Hello World!": {"English", ""},
		}

		for want, args := range testSuites {
			got, _ := Hello(args[0], args[1])
			assert(t, got, want)
		}
	})

	t.Run("Use English when no language is given", func(t *testing.T) {
		testSuites := map[string][]string{
			"Hello Thang!": {"", "Thang"},
		}

		for want, args := range testSuites {
			got, _ := Hello(args[0], args[1])
			assert(t, got, want)
		}
	})
}

package hello_world

import "fmt"

type HelloFunc = func(n string) (string, error)

var TranslateHello = map[string]string{
	"English": "Hello",
	"French":  "Bonjour",
}

func Hello(language, name string) (string, error) {
	if name == "" {
		name = "World"
	}

	if language == "" {
		return fmt.Sprintf("%s %s!", TranslateHello["English"], name), nil
	}

	if hi, ok := TranslateHello[language]; ok {
		return fmt.Sprintf("%s %s!", hi, name), nil
	}

	return "", fmt.Errorf("%s language is not supported", language)
}

func HelloIn(language string) HelloFunc {
	return func(name string) (string, error) {
		return Hello(language, name)
	}
}

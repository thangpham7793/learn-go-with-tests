package hello_world

import "fmt"

type HelloFunc = func(n string) string

var TranslateHello = map[string]string{
	"English": "Hello",
	"French":  "Bonjour",
}

func Hello(language, name string) string {
	if name == "" {
		name = "World"
	}

	if hi, ok := TranslateHello[language]; ok {
		return fmt.Sprintf("%s %s!", hi, name)
	}

	return fmt.Sprintf("%s %s!", TranslateHello["English"], name)
}

func HelloIn(language string) HelloFunc {
	return func(name string) string {
		return Hello(language, name)
	}
}

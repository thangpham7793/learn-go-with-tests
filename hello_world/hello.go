package hello_world

import "fmt"

type HelloFunc = func(n string) string

func Hello(greeting string, name string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s %s!", greeting, name)
}

func HelloWith(greeting string) HelloFunc {
	return func(name string) string {
		return Hello(greeting, name)
	}
}

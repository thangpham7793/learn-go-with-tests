package iteration

import "strings"

func Repeat(str string) (string, error) {
	return strings.Repeat(str, 3), nil
}

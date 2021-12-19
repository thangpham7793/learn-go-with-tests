package iteration

import "strings"

func Repeat(str string, time int) (string, error) {
	return strings.Repeat(str, time), nil
}

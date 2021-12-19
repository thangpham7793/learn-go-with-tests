package iteration

import (
	"fmt"
	"strconv"
	"strings"
)

type InvalidRepeatTimeError int

func (e InvalidRepeatTimeError) Error() string {
	return fmt.Sprintf("repeat: invalid repeat time: %v", strconv.Itoa(int(e)))
}

func SafeRepeat(str string, time int) (string, error) {
	if time < 1 {
		return "", InvalidRepeatTimeError(time)
	}

	return strings.Repeat(str, time), nil
}

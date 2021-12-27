package quiz_app

import (
	"fmt"
	"io"
)

func displayResult(w io.Writer, correct, total int) {
	fmt.Fprintf(w, "You answered %d out of %d questions", correct, total)
}

package quiz_app

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Question struct {
	prompt string
	answer string
}

type QuestionParser func(f io.Reader) ([]Question, error)

type UnsupportedFormatError struct {
	format string
}

func (e UnsupportedFormatError) Error() string {
	return fmt.Sprintf("parse: unsupported format %q", e.format)
}

var MissingPromptOrAnswerError = errors.New("parse: missing prompt or answer")
var NoQuestionsGivenError = errors.New("parse: no questions given")
var parsers = map[string]QuestionParser{
	"csv": parseCsv,
}

func NewQuestionParser(f *os.File) (QuestionParser, error) {
	name := f.Name()
	nameParts := strings.Split(name, ".")

	format := nameParts[len(nameParts)-1]

	if p, ok := parsers[format]; !ok {
		return nil, UnsupportedFormatError{format}
	} else {
		return p, nil
	}

}

func parseCsv(f io.Reader) ([]Question, error) {
	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		return nil, err
	}

	if len(lines) == 0 {
		return nil, NoQuestionsGivenError
	}

	questions := make([]Question, 0, len(lines))
	for _, q := range lines {
		if len(q) != 2 || q[0] == "" || q[1] == "" {
			return nil, MissingPromptOrAnswerError
		}

		q := Question{strings.TrimSpace(q[0]), strings.TrimSpace(q[1])}

		questions = append(questions, q)
	}

	return questions, nil
}

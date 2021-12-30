package quiz_app

import (
	"encoding/csv"
	"errors"
	"io"
	"strings"
)

type Question struct {
	prompt string
	answer string
}

var MissingPromptOrAnswerError = errors.New("parse: missing prompt or answer")
var NoQuestionsGivenError = errors.New("parse: no questions given")

func parse(f io.Reader) ([]Question, error) {
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

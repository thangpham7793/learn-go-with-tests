package quiz_app

import (
	"encoding/csv"
	"io"
	"strings"
)

type Question struct {
	prompt string
	answer string
}

func parse(f io.Reader) ([]Question, error) {
	content, err := csv.NewReader(f).ReadAll()

	if err != nil {
		return nil, err
	}

	questions := make([]Question, 0, len(content))
	for _, q := range content {
		new := Question{strings.TrimSpace(q[0]), strings.TrimSpace(q[1])}
		questions = append(questions, new)
	}

	return questions, nil
}

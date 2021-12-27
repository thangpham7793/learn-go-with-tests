package quiz_app

import (
	"bufio"
	"fmt"
	"io"
)

func quiz(questions []Question, input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)
	attempt := ""
	correct := 0

	for _, q := range questions {
		fmt.Println(q.prompt)
		if scanner.Scan() {
			attempt = scanner.Text()
			if attempt == q.answer {
				correct += 1
			}
		}
	}

	return correct, nil
}

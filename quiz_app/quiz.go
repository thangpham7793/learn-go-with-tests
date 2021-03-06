package quiz_app

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

func quiz(questions []Question, input io.Reader, duration time.Duration) (int, error) {
	scanner := bufio.NewScanner(input)
	attempt := ""
	correct := 0

	fmt.Println("Press enter to start")
	for {
		if scanner.Scan() {
			break
		}
	}
	done := make(chan bool)

	go func() {
		for _, q := range questions {
			fmt.Println(q.prompt)
			if scanner.Scan() {
				attempt = strings.TrimSpace(scanner.Text())
				if attempt == q.answer {
					correct += 1
				}
			}
		}

		done <- true
	}()

	select {
	case <-time.After(duration):
		return correct, nil
	case <-done:
		return correct, nil
	}
}

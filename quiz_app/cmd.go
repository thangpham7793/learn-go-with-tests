package quiz_app

import "os"

func Quiz(filePath string) {
	f, _ := os.Open(filePath)
	questions, _ := parse(f)
	correct, _ := quiz(questions, os.Stdin)
	displayResult(os.Stdout, correct, len(questions))
}

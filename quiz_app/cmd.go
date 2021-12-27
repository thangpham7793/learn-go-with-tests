package quiz_app

import (
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Run(filePath string) {
	f, err := os.Open(filePath)
	checkError(err)

	questions, err := parse(f)
	checkError(err)

	correct, err := quiz(questions, os.Stdin)
	checkError(err)

	displayResult(os.Stdout, correct, len(questions))
}

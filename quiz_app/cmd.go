package quiz_app

import (
	"flag"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Run() {
	quizFilePath := flag.String("p", "", "path to question file")
	flag.Parse()

	f, err := os.Open(*quizFilePath)
	checkError(err)

	questions, err := parse(f)
	checkError(err)

	correct, err := quiz(questions, os.Stdin)
	checkError(err)

	displayResult(os.Stdout, correct, len(questions))
}

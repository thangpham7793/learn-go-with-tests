package quiz_app

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Run() {
	quizFilePath := flag.String("p", "", "path to question file")
	duration := flag.Int64("d", 30, "quiz time in seconds")
	flag.Parse()

	f, err := os.Open(*quizFilePath)
	checkError(err)

	questions, err := parse(f)
	checkError(err)

	durationInSeconds := time.Duration(*duration) * time.Second

	correct, err := quiz(questions, os.Stdin, durationInSeconds)
	checkError(err)

	fmt.Printf("%v has passed\n", durationInSeconds)
	displayResult(os.Stdout, correct, len(questions))
}

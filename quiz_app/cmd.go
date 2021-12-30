package quiz_app

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func shuffle(qs []Question) []Question {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(qs), func(i, j int) {
		qs[i], qs[j] = qs[j], qs[i]
	})

	return qs
}

func Run() {
	quizFilePath := flag.String("p", "", "path to question file")
	duration := flag.Int64("d", 30, "quiz time in seconds")
	shuffleMode := flag.Bool("s", false, "shuffle mode or not")

	flag.Parse()

	f, err := os.Open(*quizFilePath)
	checkError(err)

	parser, err := NewQuestionParser(f)
	checkError(err)

	questions, err := parser(f)
	checkError(err)

	durationInSeconds := time.Duration(*duration) * time.Second
	if *shuffleMode {
		questions = shuffle(questions)
	}

	correct, err := quiz(questions, os.Stdin, durationInSeconds)
	checkError(err)

	fmt.Printf("\n%v has passed\n", durationInSeconds)
	displayResult(os.Stdout, correct, len(questions))
}

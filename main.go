package main

import (
	"flag"

	"github.com/thangpham7793/learn-go-with-tests/quiz_app"
)

func main() {
	quizFilePath := flag.String("p", "", "path to question file")

	flag.Parse()

	quiz_app.Run(*quizFilePath)
}

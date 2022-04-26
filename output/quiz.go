package output

import (
	"fmt"
	"github.com/igienator/quiz_game/input"
)

var QuizQuestions = input.QuizCsvFileReader()

func QuizCorrectAnswers() (int, int) {
	CorrectAnswers := QuizQuestionsOutput()
	PointLimit := len(QuizQuestions)

	return CorrectAnswers, PointLimit
}

func QuizQuestionsOutput() int {

	pointsCounter := 0
	message := "Your Answer: "

	for i := 0; i <= (len(QuizQuestions) - 1); i++ {
		fmt.Printf("Question: %v\n", QuizQuestions[i].Question)
		answer, err := input.GetInput(message)
		if err != nil {
			fmt.Println("Error reading input")
			return 0
		}

		if answer == QuizQuestions[i].Answer {
			pointsCounter++

		}

	}

	return pointsCounter
}

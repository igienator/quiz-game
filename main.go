package main

import (
	"fmt"
	"github.com/igienator/quiz_game/output"
)

var points int = 0

func main() {
	startGame()

	endGame()

}

func startGame() {
	fmt.Println("Welcome into my QUIZ")
	fmt.Println("You have 30 seconds to answer every question")

	correctAnswers, pointLimit := output.QuizCorrectAnswers()

	fmt.Printf("You have answered %v out of %v questions correctly\n", correctAnswers, pointLimit)

}

func endGame() {
	fmt.Println("END GAME")
}

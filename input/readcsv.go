package input

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Questions struct {
	Question string
	Answer   string
}

func QuizCsvFileReader() []Questions {
	csvFile, err := os.Open("problems.csv")

	if err != nil {
		fmt.Println("Error opening questions file!")
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	var questionsFromCsv []Questions

	for _, line := range csvLines {
		questionsFromCsv = append(questionsFromCsv, Questions{
			Question: line[0],
			Answer:   line[1],
		})

	}

	return questionsFromCsv
}

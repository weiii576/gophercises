package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type Question struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func loadQuestions(csvFilename *string) []Question {
	var questions []Question

	file, err := os.Open("problems.csv")
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	for _, record := range records {
		// answerInt, _ := strconv.Atoi(record[1])
		data := Question{
			question: record[0],
			answer:   record[1],
		}

		questions = append(questions, data)
	}

	return questions
}

func askQuestion(question Question, inputCh chan string) {
	fmt.Print(question.question, "=")
	input := ""
	fmt.Scanf("%s", &input)
	inputCh <- input
}

func main() {
	correctCount := 0
	inputCh := make(chan string)
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")

	questions := loadQuestions(csvFilename)

	for _, question := range questions {
		go askQuestion(question, inputCh)

		select {
		case <-time.After(10 * time.Second):
			fmt.Println("\ntimes up")
			return
		case input := <-inputCh:
			fmt.Print("\033[H\033[2J")
			if input == question.answer {
				correctCount++
			}
		}

	}

	fmt.Println("Correct:", correctCount, "/ 12")
}

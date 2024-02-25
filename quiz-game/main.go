package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type Question struct {
	question string
	answer   string
}

func loadQuestions() ([]Question, error) {
	var questions []Question

	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()

	if err != nil {
		return questions, err
	}

	for _, record := range records {
		// answerInt, _ := strconv.Atoi(record[1])
		data := Question{
			question: record[0],
			answer:   record[1],
		}

		questions = append(questions, data)
	}

	return questions, nil
}

func askQuestion(question Question, inputCh chan string) {
	fmt.Print(question.question, "=")
	input := ""
	fmt.Scanf("%s", &input)
	inputCh <- input
}

func main() {
	correctCount := 0

	questions, err := loadQuestions()

	if err != nil {
		fmt.Println(err)
	}

	timer := time.NewTimer(5 * time.Second)
	for _, question := range questions {
		timer.Reset(5 * time.Second)
		inputCh := make(chan string)

		go askQuestion(question, inputCh)

		select {
		case <-timer.C:
			fmt.Println("\ntimes up")
			return
		case input := <-inputCh:
			if input == question.answer {
				correctCount++
			}
		}

	}

	fmt.Println("Correct:", correctCount, "/ 12")
}

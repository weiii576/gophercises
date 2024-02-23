package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Question struct {
	question string
	answer   string
}

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Please check if there exist problems.csv in the same folder.")
		return
	}

	var questions []Question
	for _, record := range records {
		// answerInt, _ := strconv.Atoi(record[1])
		data := Question{
			question: record[0],
			answer:   record[1],
		}

		questions = append(questions, data)
	}
}

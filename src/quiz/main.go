package main

import (
	"encoding/csv"
	_ "flag"
	"fmt"
	"log"
	"os"
	_ "time"
)

type Quiz struct {
	Question string
	Answer   string
}

func askQuestion(q Quiz) bool {
	var userAnswer string
	fmt.Printf("Question: %s\n Your Answer: ", q.Question)
	fmt.Scanln(&userAnswer)
	return userAnswer == q.Answer
}

func createQuizList(data [][]string) []Quiz {
	var quizList []Quiz
	for i, line := range data {
		if i >= 0 {
			var rec Quiz
			for j, field := range line {
				if j == 0 {
					rec.Question = field
				} else if j == 1 {
					rec.Answer = field
				}
			}
			quizList = append(quizList, rec)
		}
	}
	return quizList
}

func main() {
	// Part 1
	if len(os.Args) < 2 {
		fmt.Println("How to run quiz:\n\tQuizt [problems CSV file]")
		return
	}

	// Open file
	path := "./src/quiz/internal/" + os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read csv file
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	quizList := createQuizList(data)

	// fmt.Printf("%+v\n", quizList)
	// Ask user the questions from the files
	// Compare user answer with file answer
	// Keep track of how many questions got correct
	var score int
	for _, quiz := range quizList {
		if askQuestion(quiz) {
			score++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", score, len(quizList))

	// Add a timer
	// Default time limit 30 sec
	// Customizable with a flag
	// Request key press to start the quiz and countdown
	// Quiz stops and outputs the results after the time limit exceeds
}

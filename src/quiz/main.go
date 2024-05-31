package main

import (
	"encoding/csv"
	_ "flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Quiz struct {
	Question string
	Answer   string
}

func quizTimer() {
	// Customizable with a flag
	// Request key press to start the quiz and countdown
	// Quiz stops and outputs the results after the time limit exceeds
}

func askQuestion(q Quiz, answerCh chan<- bool) {
	var userAnswer string
	fmt.Printf("Question: %s\n Your Answer: ", q.Question)
	fmt.Scanln(&userAnswer)
	answerCh <- (userAnswer == q.Answer)
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

	// Start the Quiz
	var start string
	fmt.Println("Start Quiz? [Y/N]")
	fmt.Scanln(&start)
	if start == "Y" || start == "y" {
		// Add a timer
		// Default time limit 30 sec
		timer := time.NewTimer(30 * time.Second)
		fmt.Println("Time started!")

		// Channel to receive answers
		answerCh := make(chan bool)
		var score int

		// Ask user the questions from the files
		// Compare user answer with file answer
		// Keep track of how many questions got correct
		// Goroutine to ask questions
		go func() {
			for _, quiz := range quizList {
				askQuestion(quiz, answerCh)
			}
			close(answerCh)
		}()

	loop:
		for {
			select {
			case <-timer.C:
				fmt.Println("Time up!")
				break loop
			case correct, ok := <-answerCh:
				if !ok {
					break loop
				}
				if correct {
					score++
				}
			}
		}
		fmt.Printf("You scored %d out of %d.\n", score, len(quizList))
	} else {
		fmt.Println("Until next time!")
	}
}

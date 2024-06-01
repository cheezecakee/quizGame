package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Quiz struct {
	Question string
	Answer   string
}

func shuffleQuiz(quizList []Quiz) {
	rand.Shuffle(len(quizList), func(i, j int) {
		quizList[i], quizList[j] = quizList[j], quizList[i]
	})
}

func askQuestion(q Quiz, answerCh chan<- bool) {
	var userAnswer string
	fmt.Printf("Question: %s\n Your Answer: ", q.Question)
	fmt.Scanln(&userAnswer)
	trimmedAns := strings.Trim(userAnswer, " ")
	answerCh <- (strings.ToLower(trimmedAns) == strings.ToLower(q.Answer))
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
					rec.Answer = strings.TrimSpace(field)
				}
			}
			quizList = append(quizList, rec)
		}
	}
	return quizList
}

func main() {
	// Customizable time limit with a flag
	var timeLimit int
	flag.IntVar(&timeLimit, "time", 30, "Time limit for the quiz in seconds")
	// Customizable shuffle the quiz order
	var shuffle bool
	flag.BoolVar(&shuffle, "shuffle", false, "Shuffles the quiz order")

	flag.Parse()

	fmt.Println("Time limit set to:", timeLimit)
	fmt.Println("Shuffle set to:", shuffle)

	if len(flag.Args()) < 1 {
		fmt.Println("How to run quiz:\n\tgo run src/quiz/main.go -time=<seconds>[optional] -shuffle[optional] -capital[optional] [CSV file]")
		return
	}

	// Open file
	path := "./src/quiz/internal/" + flag.Args()[0]

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

	// Create quiz
	quizList := createQuizList(data)

	// Shuffle quiz if set to true, default false
	if shuffle {
		shuffleQuiz(quizList)
	}

	// Start the Quiz
	var start string
	fmt.Println("Start Quiz? [Y/N]")
	fmt.Scanln(&start)
	if start == "Y" || start == "y" {
		// Default time limit 30 sec
		timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
		fmt.Println("Time started!")

		// Channel to receive answers
		answerCh := make(chan bool)
		var score int

		// Goroutine to ask questions
		go func() {
			if shuffle {
			}
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

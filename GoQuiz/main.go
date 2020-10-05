package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A CSV file in the format of {question, string}")
	timeLimit := flag.Int("time", 30, "The Time Limit for the Quiz in Seconds")

	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open file: %v", *csvFilename))
	}

	csvReader := csv.NewReader(file)

	lines, err := csvReader.ReadAll()

	if err != nil {
		exit("Could not read the questions from the given CSV")
	}

	problems := parseLines(lines)

	score := 0

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// Should I have this in the For Loop?
	answerChannel := make(chan string)

	for i, problem := range problems {		
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nThanks for taking the Quiz. You scored: %d points \n", score)
			return
		case answer := <-answerChannel:
			if answer == problem.answer {
				score++
			}
		}
	}

	fmt.Printf("Thanks for taking the Quiz. You scored: %d points \n", score)
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{line[0], strings.TrimSpace(line[1])}
	}

	return problems
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

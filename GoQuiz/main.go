package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A CSV file in the format of {question, string}")
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
	fmt.Printf("%v\n", problems)

	
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{line[0], line[1]}
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

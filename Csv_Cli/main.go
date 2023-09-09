package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "questions.csv", "Use a csv file with content formatted as 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Error reading from file: %s", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("The specified csv file could not be parsed")
	}

	problems := parseLines(lines)
	
	quizMaster(problems)

}

type problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []problem {
	output := make([]problem, len(lines))

	for i, line := range lines {
		output[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return output
}

func quizMaster(problems []problem) {

	score := 0
	for i, problem := range problems {

		var answer string
		fmt.Printf("Question #%d: %s \n", i+1, problem.question)

		fmt.Scanf("%s\n", &answer)

		if answer == problem.answer {
			score++
		}
	}

	fmt.Printf("You scored %d out of %d\n", score, len(problems))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

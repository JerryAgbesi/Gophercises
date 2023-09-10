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
	csvFilename := flag.String("csv", "questions.csv", "Use a csv file with content formatted as 'question,answer'")
	timeLimit := flag.Int("limit", 5, "Time limit for the quiz in seconds")
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
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	score := 0
	for i, problem := range problems {
		fmt.Printf("Question #%d: %s \n", i+1, problem.question)
		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)

			answerCh <- answer

		}()

		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d\n", score, len(problems))
			return
		case answer := <-answerCh:

			if answer == problem.answer {
				score++
			}

		}

	}

	// quizMaster(problems, timer)

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

// func quizMaster(problems []problem, timer *time.Timer) {

// }

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

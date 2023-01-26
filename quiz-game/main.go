// First Go program
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

// Main function
func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format 'questions,answer'")
	flag.Parse()

	quiz(*csvFileName)
}

func read_csv(filepath string) [][]string {

	file, err := os.Open(filepath)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file %s\n", filepath))
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	return records
}

type Problem struct {
	question string
	answer   string
}

func get_problems(filepath string) []Problem {
	var problems []Problem

	records := read_csv(filepath)

	for _, record := range records {
		problems = append(problems, Problem{record[0], record[1]})
	}

	return problems
}

func (p Problem) ask_the_user() bool {
	// Ask the user the question, get their input and compare the answer

	fmt.Printf("Q: %s\n", p.question)
	var user_answer string

	fmt.Scanln(&user_answer)

	if user_answer == p.answer {
		fmt.Println("Correct!")
		return true
	} else {
		fmt.Println("Wrong!")
		return false
	}

}

func quiz(fileName string) {
	var count int
	var correct_answer_count int

	var problems []Problem = get_problems(fileName)
	var number_of_questions int = len(problems)

	for _, problem := range problems {
		count += 1
		if problem.ask_the_user() {
			correct_answer_count += 1
		}
		fmt.Printf("You've got %d questions correct out of %d and there are %d questions to go\n", correct_answer_count, count, number_of_questions-count)
	}
	fmt.Printf("Thanks for playing, you scored %d out of %d", correct_answer_count, count)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

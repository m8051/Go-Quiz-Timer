/*
Go Quiz Timer

A simple command-line quiz application written in Go.

Features:
- Reads questions and answers from a CSV file.
- Limits total quiz duration with a concurrent timer.
- Tracks user score based on correct answers.
- Prints a message when time is up and shows the final score.

Usage:
- Run the program with optional command-line flags:
  --file : Path to the CSV file containing questions (default: problems.csv)
  --limit: Time limit for the quiz in seconds (default: 30)

CSV Format:
- Each row contains a question and its answer, separated by a comma.
  Example:
    5+5,10
    7+3,10
    2+2,4

Author: m8051
Date: 17/12/2025
License: MIT
*/

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

// cmdHelper parses command-line flags and validates them.
// Returns pointers to the limit and CSV file path, or an error if invalid.

func cmdHelper() (*int, *string, error) {
	limitFlag := flag.Int("limit", 30, "limit=30")
	csvFlag := flag.String("file", "problems.csv", "file=problems.csv")

	flag.Parse()

	if *limitFlag == 0 || *csvFlag == "" {
		return nil, nil, fmt.Errorf("\n\nUsage: \n\t--limit=30 \n\t--file=problems.csv")
	}

	return limitFlag, csvFlag, nil
}

// readCSV opens and reads the CSV file specified by csvFile.
// Returns a slice of records (each record is a slice of strings) or an error.

func readCSV(csvFile *string) ([][]string, error) {
	file, err := os.Open(*csvFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// calculateDuration sleeps for the specified duration (in seconds)
// and sends a signal to the channel when time is up.

func calculateDuration(duration int, ch chan int) {
	time.Sleep(time.Duration(duration) * time.Second)
	// Sends the result back to the channel
	ch <- duration
}

func main() {
	limit, csv, err := cmdHelper()
	if err != nil {
		fmt.Println("Error to parse the command line arguments:", err)
		return
	}

	var input string  // Stores user input for each question
	var score int = 0 // Tracks the number of correct answers

	records, err := readCSV(csv)
	if err != nil {
		fmt.Println("Error to read the csv file:", err)
		return
	}

	/*
		records is a type of
			[][]string{
				{"5+5", "10"},
				{"7+3", "10"},
				{"2+2", "4"},
			}
		That's why row[0] = 5+5 and row[1] = 10 in each iteration
	*/

	// Creates a channel
	ch := make(chan int)

	// runs concurrently the function calculateDuration
	go calculateDuration(*limit, ch)

Loop:
	for _, row := range records {
		fmt.Print(row[0], " = ")

		select {
		case result := <-ch:
			fmt.Println("\nTime is up after", result, "seconds")
			break Loop
		default:
			fmt.Scanln(&input)

			if row[1] == input {
				score++
			}
		}
	}

	fmt.Printf("You scored %d out of %d\n", score, len(records))
}

# Go Quiz Timer

A simple command-line quiz application written in Go. This program reads questions and answers from a CSV file, runs a timed quiz, and calculates the user's score.

---

## Features

- Reads questions and answers from a CSV file.
- Limits total quiz duration with a concurrent timer.
- Tracks user score based on correct answers.
- Prints a message when time is up and shows the final score.
- Demonstrates basic concurrency in Go using goroutines and channels.

---

## Requirements

- Go 1.18 or higher
- A CSV file containing quiz questions and answers

---

## CSV Format

The CSV file should have two columns (no header row required):

```csv
question,answer
5+5,10
7+3,10
2+2,4
```

- The first column is the question.
- The second column is the correct answer.

---

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-quiz-timer.git
cd go-quiz-timer
```

2. Run the program:

```bash
go run main.go --file=problems.csv --limit=10
```

---

## Usage

Command-line Flags
--file : Path to the CSV file containing quiz questions (default: problems.csv)
--limit : Time limit for the quiz in seconds (default: 30)


```bash
go run main.go --file=problems.csv --limit=60
```

## How It Works
1. The program parses command-line flags using the flag package.
2. Reads the CSV file with questions and answers using the encoding/csv package.
3. Starts a concurrent goroutine to handle the quiz timer.
4. Iterates through each question:
  - Prompts the user for an answer.
  - Increments the score if the answer is correct.
  - Stops the quiz immediately if the timer expires.
  - Prints the user's total score at the end of the quiz.

## Notes

- The timer runs concurrently but does not preempt user input. If the user is entering an answer when time expires, they may be allowed to finish that answer.
- This project is designed for educational purposes to demonstrate concurrency in Go.

## Author
m8051

## Date
17/12/2025

## License
MIT

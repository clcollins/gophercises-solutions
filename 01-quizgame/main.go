// This is a solution to https://github.com/gophercises/quiz
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type quizErrorMsg string

const (
	// csvOpenFailed indicates there was an error opening the CSV file
	csvOpenFailed quizErrorMsg = "Failed to open the CSV file"

	// csvReadFailed indicates a failure
	csvReadFailed quizErrorMsg = "An error occured reading the CSV file"
)

var csvFile string

func init() {
	flag.StringVar(&csvFile, "csv", "problems.csv", "Name of a CSV file to load for questions")
	flag.Parse()
}

func main() {

	input, err := os.Open(csvFile)
	defer input.Close()

	if err != nil {
		log.Fatal(csvOpenFailed)
	}

	reader := csv.NewReader(input)

	for {
		i, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%s: %s", csvReadFailed, err)
			break
		}
		fmt.Println(i[0], i[1])
	}
}
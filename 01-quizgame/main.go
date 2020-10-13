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

	// csv.NewReader implements the buffio.Reader
	reader := csv.NewReader(input)

	// iterate over the contents of the file
	for {
		i, err := reader.Read()
		// end of the file returns an io.EOF error, so handle this
		if err == io.EOF {
			break
		}
		// any other error is a problem
		if err != nil {
			log.Fatalf("%s: %s", csvReadFailed, err)
		}

		// we know our csv is in "question/answer" pairs per line
		fmt.Println(i[0], i[1])
	}
}
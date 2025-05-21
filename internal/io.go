package internal

import (
	"log"
	"os"
)

// ReadInputFile reads the input file and returns its contents.
func ReadInputFile() string {
	file, openErr := os.Open(InputFileName)

	if openErr != nil {
		log.Fatal(openErr)
	}

	defer file.Close()

	data, readErr := os.ReadFile(InputFileName)

	if readErr != nil {
		log.Fatal(readErr)
	}

	return string(data)
}

// WriteOutputFile writes the output file.
func WriteOutputFile(output string) {
	err := os.WriteFile(OutputFileName, []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

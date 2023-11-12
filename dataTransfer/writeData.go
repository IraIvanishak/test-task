package dataTransfer

import (
	"encoding/json"
	"log"
	"os"

	"github.com/IraIvanishak/test-task-contentquo/breeds"
)

// WriteResultsToFile writes the formatted breeds data to a JSON file.
func WriteResultsToFile(formattedBreeds breeds.FormattedBreeds) {
	file, err := os.Create("out.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(formattedBreeds)
	if err != nil {
		log.Fatalf("Error encoding to JSON: %v", err)
	}

	log.Println("Data successfully written to the file out.json.")
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/yeboahnanaosei/gitplus/csv"
)

func main() {
	csvFilepath := flag.String("file", "", "The path to the file the CSV file")
	flag.Parse()

	if *csvFilepath == "" {
		fmt.Fprintf(os.Stderr, "csvchecker: %v\n", "You did not supply the path to the file. Please supply '-file /path/to/csv/file'")
		os.Exit(1)
	}

	csvFile, err := os.Open(*csvFilepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "csvchecker: %v %v\n", "Failed to open supplied csv file", err)
		os.Exit(1)
	}
	defer csvFile.Close()

	validRecords, invalidRecords, err := csv.Validate(csvFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "csvchecker: %v\n", err)
		os.Exit(1)
	}

	payload := map[string]interface{}{
		"ok":             false,
		"validRecords":   nil,
		"invalidRecords": nil,
	}

	if len(invalidRecords) == 0 {
		payload["ok"] = true
		payload["validRecords"] = validRecords
	} else {
		payload["ok"] = false
		payload["invalidRecords"] = invalidRecords
		payload["validRecords"] = validRecords
	}

	json.NewEncoder(os.Stdout).Encode(payload)
}

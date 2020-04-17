package cvet

import (
	"encoding/csv"
	"io"
	"strings"
)

// invalidRecord is a record or row in the csv file that has at least
// one empty column.
type invalidRecord struct {
	RowNumber int      `json:"row"`
	Columns   []string `json:"cols"`
}

// Validate validates f as a valid csv file with valid data.
func Validate(f io.Reader) (validRecords [][]string, invalidRecords[]InvalidRecord, err error) {
	r := csv.NewReader(f)
	r.TrimLeadingSpace = true

	uploadedCSV, err := r.ReadAll()
	if err != nil {
		return validRecords, invalidRecords, err
	}

	// The first row in the csv is usually the header which has the name of each
	// column in the csv file
	var header []string = uploadedCSV[0]
	headerLength := len(header)

	// To determine the row number of an invalid row, we need to account for
	// the header in the file
	const headerOffset = 2

	// Skip the header. Go through each row in the file checking that for each
	// row there are no empty columns
	for rowIndex, record := range uploadedCSV[1:] {
		currentRecord := new(invalidRecord)
		currentRecord.RowNumber = rowIndex + headerOffset
		recordIsValid := true

		for columnIndex, field := range record {
			if strings.TrimSpace(field) == "" {
				recordIsValid = false
				currentRecord.Columns = append(currentRecord.Columns, header[columnIndex])
			}
		}

		if recordIsValid {
			validRecords = append(validRecords, record)
		} else if !recordIsValid && len(currentRecord.Columns) != headerLength {
			invalidRecords = append(invalidRecords, *currentRecord)
		}
	}
	return validRecords, invalidRecords, nil
}

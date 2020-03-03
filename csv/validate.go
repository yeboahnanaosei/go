package csv

import (
	"encoding/csv"
	"io"
	"strings"
)

// InvalidRecord is a record that is not valid.
type InvalidRecord struct {
	RowNumber int      `json:"row"`
	Columns   []string `json:"columns"`
}

// Validate validates f as a valid csv file with valid data.
func Validate(f io.Reader) (validRecords [][]string, invalidRecords[]InvalidRecord, err error) {
	// TODO: Run a check to ensure that we are getting a valid csv file
	// Possibly we can check for the mime type

	r := csv.NewReader(f)
	r.TrimLeadingSpace = true

	uploadedCSV, err := r.ReadAll()
	if err != nil {
		return validRecords, invalidRecords, err
	}

	rowOffset := 2
	header := uploadedCSV[0]
	headerLength := len(header)

	for row, record := range uploadedCSV[1:] {
		currentRecord := new(InvalidRecord)
		currentRecord.RowNumber = row + rowOffset
		recordIsValid := true

		for column, field := range record {
			if strings.Trim(field, " ") == "" {
				recordIsValid = false
				currentRecord.Columns = append(currentRecord.Columns, header[column])
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

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/yeboahnanaosei/gitplus/csv"
)

func main() {
	out := make(map[string]interface{})

	csvFilepath := flag.String("file", "", "The path to the file the CSV file")
	flag.Parse()

	if *csvFilepath == "" {
		out["status"] = "fail"
		out["code"] = 400
		out["msg"] = "Bad request"
		out["error"] = map[string]string{
			"msg": "no path to csv file supplied. a path to a csv file was expected but none was supplied",
			"fix": "use the '-file' flag to supply the path eg. -file /path/to/csv/file",
		}
		json.NewEncoder(os.Stdout).Encode(out)
		return
	}

	csvFile, err := os.Open(*csvFilepath)
	if err != nil {
		out["status"] = "fail"
		out["code"] = 500
		out["msg"] = "An internal error occured"
		out["error"] = map[string]string{
			"msg": fmt.Sprintf("there was an error trying to process the csv file. server said: %v", err),
			"fix": "please make sure you provided is a valid csv file. If this continues, please wait and try again later. You can also contact support",
		}
		json.NewEncoder(os.Stdout).Encode(out)
		return
	}
	defer csvFile.Close()

	validRecords, invalidRecords, err := csv.Validate(csvFile)

	if err != nil {
		out["status"] = "fail"
		out["code"] = 500
		out["msg"] = "An internal error occured"
		out["error"] = map[string]string{
			"msg": fmt.Sprintf("there was an error trying to process the csv file. server said: %v", err),
			"fix": "please make sure you provided is a valid csv file. If this continues, please wait and try again later. You can also contact support",
		}
		json.NewEncoder(os.Stdout).Encode(out)
		return
	}

	out["status"] = "success"
	out["code"] = 200
	out["msg"] = "operation was successful"
	out["data"] = map[string]interface{}{
		"validRecords":   validRecords,
		"invalidRecords": invalidRecords,
	}
	json.NewEncoder(os.Stdout).Encode(out)
	return
}

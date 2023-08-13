package main

/*
  This code will read in regex.csv into a slice. The getRegexValue
  function will return the regex value for a specified variable name.
  todo: read and flatten the variables in terraform.tfvars.
  todo: loop through all variables, get the regex,
		check the value against the regex. Return true or false.
*/

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type varRegex struct {
	Var   string
	Regex string
}

func main() {
	// Load the CSV file into a slice of varRegex structures
	varData := []varRegex{}
	fn := "regex.csv"
	f, err := os.Open(fn)
	if err != nil {
		//log.Fatal(err)
		fmt.Println("file error")
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if record[0] != "" && record[1] != "" && record[2] != "" {
			varData = append(varData, varRegex{
				Var:   record[0],
				Regex: record[2],
			})
		}
	}

	// Function to get the value of the "regex" column for a given "var" value
	getRegexValue := func(varName string) string {
		for _, v := range varData {
			if v.Var == varName {
				return v.Regex
			}
		}
		return ""
	}

	// Print the value of the "regex" column for "var.location"
	fmt.Println("Location regex=", getRegexValue("var.location"))
	fmt.Println("UAMI regex=", getRegexValue("var.uami-cp"))
}

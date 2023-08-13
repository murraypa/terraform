package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {

	f, err := os.Open("terraform.tfvars")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewWriter(os.Stdout)

	reader := bufio.NewReader(f)

	var lines []string
	var currentKey string

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "#") { //ignore comment lines
			continue
		}

		if strings.HasPrefix(line, "\"") {
			parts := strings.SplitN(line, "\"", 3)
			if len(parts) > 1 {
				r.Write([]string{parts[1], parts[2]})
			}
		} else if strings.HasPrefix(line, "}") {
			currentKey = ""
		} else if currentKey != "" {
			key := currentKey + "." + strings.TrimSpace(line)
			lines = append(lines, key)
		} else if line != "" {
			currentKey = line
		}
	}

	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)

		// Remove trailing "{" if present
		if strings.HasSuffix(parts[0], "{") {
			line = strings.TrimSuffix(line, "{")
		}
		fmt.Println(parts[0], "|", parts[1])

		if len(parts) > 1 {
			r.Write([]string{strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])})
		}
	}

	r.Flush()
}

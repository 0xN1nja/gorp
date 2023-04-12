package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ttacon/chalk"
)

func search(query string, file_path string) []string {
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("Unable To Read File")
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var results []string
	for scanner.Scan() {
		file_content := scanner.Text()
		for _, lines := range strings.Split(file_content, "\n") {
			if strings.Contains(lines, query) {
				results = append(results, strings.ReplaceAll(lines, query, chalk.Red.Color(query)))
			}
		}
	}
	return results
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Incorrect Number Of Arguments Supplied")
		os.Exit(1)
	}
	query := os.Args[1]
	file_path := os.Args[2]
	results := search(query, file_path)
	for _, result := range results {
		fmt.Println(result)
	}
}

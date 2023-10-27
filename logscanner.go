package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"strings"
)

func GetErrorLogs() ([]string, error) {
	file, err := os.Open("loglist.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var errorLogs []string
	for _, record := range records {
		if len(record) == 0 {
			continue
		}

		filePath := record[0]
		if containsError(filePath) {
			errorLogs = append(errorLogs, filePath)
		}
	}

	return errorLogs, nil
}

func containsError(filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "error") {
			return true
		}
	}
	return false
}

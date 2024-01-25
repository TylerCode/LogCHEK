package scanner

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// ScanLogs scans a list of log files and reports which files contain errors.
func ScanLogs(reportScanStatus func(string), reportErrors func([]string)) {
	reportScanStatus("Scanning")
	file, err := os.Open("loglist.csv")
	if err != nil {
		reportScanStatus(fmt.Sprintf("Error opening loglist.csv: %v", err))
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		reportScanStatus(fmt.Sprintf("Error reading loglist.csv: %v", err))
		return
	}

	var errorLogs []string
	for _, record := range records {
		if len(record) == 0 {
			continue
		}

		filePath := record[0]
		if ContainsError(filePath) {
			errorLogs = append(errorLogs, filePath)
		}
	}

	if len(errorLogs) == 0 {
		reportScanStatus("All Clear")
	} else {
		reportScanStatus("Errors Found")
		reportErrors(errorLogs)
	}
}

// ContainsError checks if a given log file contains the word "error".
func ContainsError(filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Scan line by line
	for scanner.Scan() {
		// Convert text to lower case
		text := strings.ToLower(scanner.Text())

		// Check if 'error' is in the text
		if errorIndex := strings.Index(text, "error"); errorIndex != -1 {
			// If 'error' is at the beginning of the text, it's an error
			if errorIndex == 0 {
				return true
			}

			// Check the text before 'error'
			beforeError := text[:errorIndex]
			trimmedBeforeError := strings.TrimSpace(beforeError)

			// Check if the last character before 'error' is '0' not preceded by another digit
			if len(trimmedBeforeError) > 0 {
				lastCharIndex := len(trimmedBeforeError) - 1
				if trimmedBeforeError[lastCharIndex] == '0' {
					if lastCharIndex == 0 || !unicode.IsDigit(rune(trimmedBeforeError[lastCharIndex-1])) {
						continue
					}
				}
			}

			// Otherwise, it's an error
			return true
		}
	}

	return false
}

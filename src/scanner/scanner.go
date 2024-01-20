package scanner

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
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

	for scanner.Scan() {
		// Convert text to lower case
		text := strings.ToLower(scanner.Text())

		// Check if 'error' is in the text
		if errorIndex := strings.Index(text, "error"); errorIndex != -1 {
			// If 'error' is at the beginning of the text, it's an error
			if errorIndex == 0 {
				return true
			}

			// If there is space before 'error', check the substring before that space
			if errorIndex > 1 && text[errorIndex-1] == ' ' {
				// Get substring before the space
				precedingText := text[:errorIndex-1]
				if !strings.HasSuffix(precedingText, "0") {
					return true
				}
			} else if errorIndex == 1 || (errorIndex > 1 && text[errorIndex-2] != '0') {
				// If there's no space before 'error', check the character before 'error'
				// Or if there are more characters before 'error', ensure the second last is not '0'
				return true
			}
		}
	}

	return false
}

package scanner

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"regexp"
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

		// Use regular expression to find "error" not preceded by "0"
		matched, _ := regexp.MatchString(`[^0]\serror`, text)
		if matched || strings.HasPrefix(text, "error") {
			return true
		}
	}

	return false
}

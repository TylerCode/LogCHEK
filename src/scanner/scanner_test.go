package scanner

import (
	"os"
	"testing"
)

func createDummyLogFile(name string, content string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

func TestScanLogs(t *testing.T) {
	// Setup: Create dummy log files and loglist.csv
	logFiles := []struct {
		name    string
		content string
	}{
		{"log1.txt", "INFO: Nothing to see here"},
		{"log2.txt", "ERROR: EVERYTHING IS ON FIRE"},
		{"log3.txt", "INFO: Process run with 0 errors."},
	}

	for _, lf := range logFiles {
		err := createDummyLogFile(lf.name, lf.content)
		if err != nil {
			t.Fatalf("Failed to create dummy log file: %v", err)
		}
		defer os.Remove(lf.name)
	}

	err := createDummyLogFile("loglist.csv", "log1.txt\nlog2.txt\nlog3.txt")
	if err != nil {
		t.Fatalf("Failed to create dummy loglist.csv: %v", err)
	}
	defer os.Remove("loglist.csv")

	// Define mock functions to capture the output
	var scanStatus string
	var errorLogs []string
	reportScanStatus := func(status string) {
		scanStatus = status
	}
	reportErrors := func(logs []string) {
		errorLogs = logs
	}

	// Run the function under test
	ScanLogs(reportScanStatus, reportErrors)

	// Assertions
	if scanStatus != "Errors Found" {
		t.Errorf("Expected 'Errors Found', got '%v'", scanStatus)
	}

	if len(errorLogs) != 1 || errorLogs[0] != "log2.txt" {
		t.Errorf("Expected ['log2.txt'], got '%v'", errorLogs)
	}
}

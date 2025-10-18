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
		{"log1.txt", "2024-01-11T20:55:42+52:33 INFO: Nothing to see here"},
		{"log2.txt", "2024-01-11T20:55:49+31:00 ErRoR: EVERYTHING IS ON FIRE"},
		{"log3.txt", "2024-01-11T20:55:56+51:05 INFO: Process run with 0 errors."},
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

func TestScanLogsWithInaccessibleFiles(t *testing.T) {
	// Setup: Create loglist.csv with a non-existent file
	err := createDummyLogFile("loglist.csv", "nonexistent.txt\n")
	if err != nil {
		t.Fatalf("Failed to create dummy loglist.csv: %v", err)
	}
	defer os.Remove("loglist.csv")

	// Define mock functions to capture the output
	var scanStatuses []string
	reportScanStatus := func(status string) {
		scanStatuses = append(scanStatuses, status)
	}
	reportErrors := func(logs []string) {
		// No error logs expected for inaccessible files
	}

	// Run the function under test
	ScanLogs(reportScanStatus, reportErrors)

	// Assertions - should report warning about inaccessible files
	foundWarning := false
	for _, status := range scanStatuses {
		if status == "Warning: 1 file(s) could not be accessed" {
			foundWarning = true
			break
		}
	}
	if !foundWarning {
		t.Errorf("Expected warning about inaccessible files, got: %v", scanStatuses)
	}

	// Should still report "All Clear" since no accessible files had errors
	if scanStatuses[len(scanStatuses)-1] != "All Clear" {
		t.Errorf("Expected final status 'All Clear', got '%v'", scanStatuses[len(scanStatuses)-1])
	}
}

func TestContainsError(t *testing.T) {
	tests := []struct {
		name             string
		content          string
		expectError      bool
		expectAccessible bool
	}{
		{"no error", "INFO: Everything is fine", false, true},
		{"has error", "ERROR: Something went wrong", true, true},
		{"has 0 errors", "Process completed with 0 errors", false, true},
		{"mixed case error", "ErRoR: Failed", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename := "test_" + tt.name + ".txt"
			err := createDummyLogFile(filename, tt.content)
			if err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}
			defer os.Remove(filename)

			hasError, accessible := ContainsError(filename)
			if accessible != tt.expectAccessible {
				t.Errorf("Expected accessible=%v, got %v", tt.expectAccessible, accessible)
			}
			if hasError != tt.expectError {
				t.Errorf("Expected hasError=%v, got %v", tt.expectError, hasError)
			}
		})
	}

	// Test inaccessible file
	t.Run("inaccessible file", func(t *testing.T) {
		hasError, accessible := ContainsError("nonexistent_file.txt")
		if accessible != false {
			t.Errorf("Expected accessible=false for nonexistent file, got %v", accessible)
		}
		if hasError != false {
			t.Errorf("Expected hasError=false for nonexistent file, got %v", hasError)
		}
	})
}

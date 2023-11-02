package main

import (
	"bufio"
	"fmt"
	"os"
	"yourModuleName/scanner" // replace with your actual module name
)

func main() {
	fmt.Println("Starting log file scan...")

	// Report functions for the CLI.
	reportScanStatus := func(status string) {
		fmt.Println(status)
	}
	reportErrors := func(errorLogs []string) {
		if len(errorLogs) > 0 {
			fmt.Println("The following log files contain errors:")
			for _, log := range errorLogs {
				fmt.Println(log)
			}
		}
	}

	// Start scanning logs.
	scanner.ScanLogs(reportScanStatus, reportErrors)

	// Wait for user input before exiting.
	fmt.Println("Press 'Enter' to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

package main

import (
	"LogCHEK/scanner"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting log file scan...")

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

	scanner.ScanLogs(reportScanStatus, reportErrors)

	fmt.Println("Press 'Enter' to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

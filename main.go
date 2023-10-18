package main

import (
	"fmt"
	"bufio"
	"encoding/csv"
	"os"
	"strings"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("LogCHEK")
	myWindow.Resize(fyne.NewSize(640, 400))  // Set window size

	statusLabel := widget.NewLabel("Ready")
	errorLogsTextArea := widget.NewMultiLineEntry()
	errorLogsTextArea.Disable()


	// Make errorLogsTextArea taller using a scroll container
	scrollContainer := container.NewVScroll(errorLogsTextArea)
	scrollContainer.SetMinSize(fyne.NewSize(640, 300))

	startButton := widget.NewButton("Start Scan", func() {
		go scanLogs(statusLabel, errorLogsTextArea)
	})

	content := container.NewVBox(
		statusLabel,
		scrollContainer,   // Use the scroll container here
		startButton,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func scanLogs(statusLabel *widget.Label, errorLogsTextArea *widget.Entry) {
	statusLabel.SetText("Scanning")
	file, err := os.Open("loglist.csv")
	if err != nil {
		statusLabel.SetText(fmt.Sprintf("Error opening loglist.csv: %v", err))
		return
	}
	defer file.Close()
	
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		statusLabel.SetText(fmt.Sprintf("Error reading loglist.csv: %v", err))
		return
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

	if len(errorLogs) == 0 {
		statusLabel.SetText("All Clear")
	} else {
		statusLabel.SetText("Errors Found")
		errorLogsTextArea.SetText(strings.Join(errorLogs, "\n"))
	}
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


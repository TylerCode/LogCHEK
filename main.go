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
	myWindow := myApp.NewWindow("Log Scanner")

	statusLabel := widget.NewLabel("Ready")
	errorLogsTextArea := widget.NewMultiLineEntry()
	startButton := widget.NewButton("Start Scan", func() {
		go scanLogs(statusLabel, errorLogsTextArea)
	})

	content := container.NewVBox(
		statusLabel,
		errorLogsTextArea,
		startButton,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func scanLogs(statusLabel *widget.Label, errorLogsTextArea *widget.Entry) {
	statusLabel.SetText("Scanning")
	errorLogs, err := GetErrorLogs()
	if err != nil {
		statusLabel.SetText("Error: " + err.Error())
		return
	}

	if len(errorLogs) == 0 {
		statusLabel.SetText("All Clear")
	} else {
		statusLabel.SetText("Errors Found")
		errorLogsTextArea.SetText(strings.Join(errorLogs, "\n"))
	}
}


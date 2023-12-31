package main

import (
    "log"
    "strings"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "LogCHEK/scanner"
)

func main() {
	myApp := app.New()

	icon, err := fyne.LoadResourceFromPath("logo.png")
	if err != nil {
		log.Printf("Failed to load icon: %v", err)
	}

	myWindow := myApp.NewWindow("LogCHEK")
	
	if icon != nil {
		myWindow.SetIcon(icon)
	}
	
	myWindow.Resize(fyne.NewSize(640, 400))

	statusLabel := widget.NewLabel("Ready")
	errorLogsTextArea := widget.NewMultiLineEntry()
	errorLogsTextArea.Disable()

	scrollContainer := container.NewVScroll(errorLogsTextArea)
	scrollContainer.SetMinSize(fyne.NewSize(640, 300))

	startButton := widget.NewButton("Start Scan", func() {
		go scanner.ScanLogs(
			func(status string) {
				statusLabel.SetText(status)
			},
			func(errorLogs []string) {
				errorLogsTextArea.SetText(strings.Join(errorLogs, "\n"))
			},
		)
	})

	content := container.NewVBox(
		statusLabel,
		scrollContainer,
		startButton,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

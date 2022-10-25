package main

import (
	"fmt"
	// "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	output *widget.Label
}
func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	var myApp App
	output, entry, btn := myApp.makeUI()

	w.SetContent(container.NewVBox(output, entry, btn))
	w.ShowAndRun()
	cleanup()
}

func (app *App) makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
	output := widget.NewLabel("Hello World")
	entry := widget.NewEntry()
	btn := widget.NewButton("Enter", func() {
		app.output.SetText(entry.Text)
	})
	btn.Importance = widget.MediumImportance

	app.output = output
	return output, entry, btn
}

func cleanup() {
	fmt.Println("cleaning stuff up")
}
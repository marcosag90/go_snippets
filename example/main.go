package main

import (
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile fyne.URI
	SaveMenuItem *fyne.MenuItem
}

func main() {
	// Create App
	a := app.New()

	// Create a window for the app
	w := a.NewWindow("Markdown Editor")

	// Get the user interface
	var cfg config
	edit, preview := cfg.makeUI()

	// set content to the window
	w.SetContent(container.NewHSplit(edit, preview))
	cfg.createMenuItems(w)

	// show window and run app
	// w.Resize(fyne.Size(800, 600))
	// w.CenterOnScreen()
	w.ShowAndRun()
}


func (app *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	app.EditWidget = edit
	app.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}

func (app *config) createMenuItems( win fyne.Window ) {
	openMenuItem := fyne.NewMenuItem("Open...", app.openFileFunc(win))
	saveMenuItem := fyne.NewMenuItem( "Save", func(){ return } )
	saveAsMenuItem := fyne.NewMenuItem( "Save as...", app.saveAsFunc(win) )
	app.SaveMenuItem = saveMenuItem
	app.SaveMenuItem.Disabled = true;
	fileMenu := fyne.NewMenu( "File", openMenuItem, saveAsMenuItem, saveMenuItem )
	menuBar := fyne.NewMainMenu( fileMenu )
	win.SetMainMenu(menuBar)
}

func (app *config) saveFunc( win fyne.Window ) func() {
	return func(){
		write, err := storage.Writer(app.CurrentFile)
		if err != nil {
			dialog.ShowError(err, win)
			return 
		}
		write.Write([]byte(app.EditWidget.Text))
		defer write.Close()
	}
}

func (app *config) saveAsFunc( win fyne.Window ) func() {
	return func(){
		callback := func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return 
			}
			if write == nil {//
				return // user cancelled
			}
			write.Write([]byte(app.EditWidget.Text))
			app.CurrentFile = write.URI()
			defer write.Close()
			
			win.SetTitle( win.Title() + " - " + write.URI().Name() )
			app.SaveMenuItem.Disabled = false
		}
		saveDialog := dialog.NewFileSave(callback , win) 

		saveDialog.Show()
	}
}

func (app *config) openFileFunc( win fyne.Window ) func() {
	return func(){
		callback := func(read fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return 
			}
			if read == nil {
				return // user cancelled
			}
			data, err := ioutil.ReadAll(read)
			if err != nil {
				dialog.ShowError(err, win)
			}
			app.EditWidget.SetText(string(data))
			app.CurrentFile = read.URI()
			defer read.Close()
			
			win.SetTitle( win.Title() + " - " + read.URI().Name() )
			app.SaveMenuItem.Disabled = false
		}
		openDialog := dialog.NewFileOpen(callback , win) 

		openDialog.Show()
	}
}
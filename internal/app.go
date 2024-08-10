package internal

import (
	"kuroneko/internal/ui"

	"fyne.io/fyne/v2/app"
)

func RunApp() {
	application := app.New()
	window := ui.NewMainWindow(application)
	window.ShowAndRun()
}

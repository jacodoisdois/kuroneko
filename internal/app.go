package internal

import (
	"kuroneko/internal/ui"

	"fyne.io/fyne/v2/app"
)

func RunApp() {
	application := app.NewWithID("com.github.jacodoisdois.kuroneko")
	application.Settings().SetTheme(&ui.OffWhiteTheme{})
	window := ui.NewMainWindow(application)
	window.ShowAndRun()
}

package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ThrowError(w fyne.Window, message string) {
	content := widget.NewLabel(message)
	customDialog := dialog.NewCustom("Error", "OK", content, w)

	customDialog.Resize(fyne.NewSize(100, 100))
	customDialog.Show()
}

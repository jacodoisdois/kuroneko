package ui

import (
	"kuroneko/internal/handler"
	errorUtils "kuroneko/internal/ui/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateTitleContainer() fyne.CanvasObject {
	title := widget.NewLabel("Kuroneko")
	title.Alignment = fyne.TextAlignCenter

	image := canvas.NewImageFromFile("assets/images/logo.jpg")
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(50, 50))
	image.Resize(fyne.NewSize(50, 50))

	spacer := layout.NewSpacer()

	titleContainer := container.NewCenter(
		container.NewHBox(
			image,
			spacer,
			title,
		),
	)

	return titleContainer
}

func CreateInputContainer(w fyne.Window) fyne.CanvasObject {
	folderLabel := widget.NewLabel("No folder selected")
	folderLabel.Wrapping = fyne.TextWrap(fyne.TextTruncateClip)

	background := canvas.NewRectangle(theme.InputBackgroundColor())
	background.SetMinSize(fyne.NewSize(200, 40))

	labelContainer := container.NewStack(
		background,
		container.NewPadded(folderLabel),
	)

	folderButton := widget.NewButton("Browse", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if uri != nil {
				folderLabel.SetText(uri.Path())
			}
		}, w)
	})

	runButton := widget.NewButton("Run", func() {
		if folderLabel.Text == "No folder selected" || folderLabel.Text == "" {
			errorUtils.ThrowError(w, "No folder selected. Please select a folder.")
		} else {
			handler.GenerateCbzFiles(folderLabel.Text, w)
		}
	})

	inputContainer := container.NewHBox(labelContainer, folderButton, runButton)

	return inputContainer
}

func NewMainWindow(a fyne.App) fyne.Window {
	window := a.NewWindow("Kuroneko")
	window.Resize(fyne.NewSize(600, 600))

	titleContainer := CreateTitleContainer()
	inputContainer := CreateInputContainer(window)
	mainWindowSpacer := canvas.NewRectangle(nil)
	mainWindowSpacer.SetMinSize(fyne.NewSize(0, 20))
	content := container.NewCenter(container.NewVBox(titleContainer, mainWindowSpacer, inputContainer))
	window.SetContent(content)

	return window
}
